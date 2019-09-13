package resources


import (
	"fmt"
	"net/http"
)

func WelcomeToAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Person API with MongoDB")

}
