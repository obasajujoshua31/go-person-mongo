package repository

import (
	"context"
	"go-mongo/channels"
	"go-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PeopleRepository interface {
	Insert(models.Person) (*mongo.InsertOneResult, error)
	FindAll() ([]*models.Person, error)
	FindByID(primitive.ObjectID) (*models.Person, error)
}

type repositoryPerson struct {
	ctx      context.Context
	database *mongo.Database
}

func NewPersonRepository(ctx context.Context, database *mongo.Database) *repositoryPerson {
	return &repositoryPerson{
		ctx:      ctx,
		database: database,
	}
}

func (d *repositoryPerson) Insert(person models.Person) (result *mongo.InsertOneResult, err error) {
	done := make(chan bool)
	go func(ch chan<- bool) {
		collection := d.database.Collection("people")
		result, err = collection.InsertOne(d.ctx, person)
		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return result, nil
	}
	return nil, err
}

func (d *repositoryPerson) FindAll() (persons []*models.Person, err error) {
	done := make(chan bool)
	go func(ch chan<- bool) {
		collection := d.database.Collection("people")
		cur, err := collection.Find(d.ctx, bson.D{})
		defer cur.Close(d.ctx)

		for cur.Next(d.ctx) {
			var person models.Person
			err := cur.Decode(&person)
			if err != nil {
				ch <- false
			}
			persons = append(persons, &person)
		}
		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return persons, nil
	}
	return nil, err
}

func (d *repositoryPerson) FindByID(id primitive.ObjectID) (person *models.Person, err error) {
	done := make(chan bool)
	filter := bson.M{"_id": id}
	go func(ch chan<- bool) {
		collection := d.database.Collection("people")
		err := collection.FindOne(d.ctx, filter).Decode(&person)
		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return person, nil
	}
	return nil, err
}
