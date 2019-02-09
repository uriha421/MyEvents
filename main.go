package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type eventServiceHandler struct{}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}

// this function receives an endpoint such as ":8181"
// and start a web server with a router
func serveAPI(endpoint string) error {
	// create a new router
	r := mux.NewRouter()
	eventsrouter = r.PathPrefix("/events").Subrouter()

	// get an event with the received ID or Name
	eventsrouter.Methods("GET").Path("/{SerachCriteria}/{search}").HandlerFunc(handler.findEventHandler)

	// get all events
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

	// register a new event
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	// start a web server
	return http.ListenAndServe(endpoint, r)
}

func main() {
	fmt.Println("main")
}
