package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hbd/go-crm/people"
	"github.com/sirupsen/logrus"
)

/* Test data.
people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Status: "Prospect", Address: &Address{City: "City X", State: "State X"}})
people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
*/

func main() {
	router := mux.NewRouter()

	c, err := people.NewController()
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to initialize controller.")
	}

	router.HandleFunc(
		"/healthcheck",
		c.Healthcheck).
		Methods(http.MethodGet)

	router.HandleFunc(
		"/people",
		c.GetPeople).
		Methods(http.MethodGet)

	router.HandleFunc(
		"/people/{"+people.PathParamID+"}",
		c.GetPerson).
		Methods(http.MethodGet)

	router.HandleFunc(
		"/people/{"+people.PathParamID+"}/status",
		c.GetStatus).
		Methods(http.MethodGet)

	router.HandleFunc(
		"/people/{"+people.PathParamID+"}/status/{"+people.PathParamNewStatus+"}",
		c.SetStatus).
		Methods(http.MethodPut)

	router.HandleFunc("/people/{"+people.PathParamID+"}",
		c.CreatePerson).
		Methods(http.MethodPost)

	router.HandleFunc("/people/{"+people.PathParamID+"}",
		c.DeletePerson).
		Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", router))
}
