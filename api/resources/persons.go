package resources

import (
	"context"
	"encoding/json"
	"errors"
	"go-mongo/database"
	"go-mongo/models"
	"go-mongo/repository"
	"go-mongo/responses"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var person models.Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	database := database.ConnectToDatabase()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	repo := repository.NewPersonRepository(ctx, database)

	func(personRepository repository.PeopleRepository) {

		result, err := personRepository.Insert(person)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		oid := result.InsertedID.(primitive.ObjectID)

		person.ID = &oid
		responses.JSON(w, http.StatusCreated, person)
	}(repo)

}

func GetPersons(w http.ResponseWriter, r *http.Request) {

	database := database.ConnectToDatabase()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	repo := repository.NewPersonRepository(ctx, database)

	func(personRepository repository.PeopleRepository) {

		result, err := personRepository.FindAll()

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, result)
	}(repo)

}

func GetOnePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	database := database.ConnectToDatabase()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	repo := repository.NewPersonRepository(ctx, database)

	func(personRepository repository.PeopleRepository) {

		result, err := personRepository.FindByID(id)

		if result == nil {
			responses.ERROR(w, http.StatusNotFound, errors.New("person is not found"))
			return
		}

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, result)
	}(repo)

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {

	// Get params
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var person models.Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	database := database.ConnectToDatabase()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	repo := repository.NewPersonRepository(ctx, database)

	func(personRepository repository.PeopleRepository) {

		result, err := personRepository.Update(id, person)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, map[string]int64{"updatedCount": result.ModifiedCount})
	}(repo)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

	// Get params
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	database := database.ConnectToDatabase()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	repo := repository.NewPersonRepository(ctx, database)

	func(personRepository repository.PeopleRepository) {

		result, err := personRepository.Delete(id)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusNoContent, map[string]int64{"deletedCount": result.DeletedCount})
	}(repo)
}
