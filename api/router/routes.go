package router

import (
	resources "go-mongo/api/resources"
	"net/http"
)

var routes = []Route{
	Route{
		URI:     "/persons",
		Method:  http.MethodPost,
		Handler: resources.CreatePerson,
	},
	Route{
		URI:     "/",
		Method:  http.MethodGet,
		Handler: resources.WelcomeToAPI,
	},
	Route{
		URI:     "/persons",
		Method:  http.MethodGet,
		Handler: resources.GetPersons,
	},
	Route{
		URI:     "/persons/{id}",
		Method:  http.MethodGet,
		Handler: resources.GetOnePerson,
	},
	Route{
		URI:     "/persons/{id}",
		Method:  http.MethodPut,
		Handler: resources.UpdatePerson,
	},
	Route{
		URI:     "/persons/{id}",
		Method:  http.MethodDelete,
		Handler: resources.DeletePerson,
	},
}
