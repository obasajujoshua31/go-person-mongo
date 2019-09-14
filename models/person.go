package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string              `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string              `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func (p Person) Validate(action string) error {
	switch action {
	case "create":
		if p.Firstname == "" {
			return errors.New("Firstname required")
		}
		if p.Lastname == "" {
			return errors.New("Lastname required")
		}
		return nil
	case "update":
		if p.Firstname == "" {
			return errors.New("Firstname required")
		}
		if p.Lastname == "" {
			return errors.New("Lastname required")
		}
		return nil
	default:
		{
			return nil
		}
	}
}
