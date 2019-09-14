package database

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnectToDatabase(t *testing.T) {
	tests := []struct {
		name string
		want *mongo.Database
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectToDatabase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectToDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}
