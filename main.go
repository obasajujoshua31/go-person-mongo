package main

import (
	"fmt"
	"log"
	"net/http"

	"go-mongo/api/router"

	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
)

func main() {
	fmt.Println("Go application started")
	HandleRequest()
}

func HandleRequest() {
	r := mux.NewRouter().StrictSlash(true)
	router.SetUpRoutes(r)
	fmt.Println("Server started at port 3200...")
	log.Fatal(http.ListenAndServe(":3200", httplogger.Golog(r)))

}
