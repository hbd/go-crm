package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/benjamin28/go-crm-fork/handlers"
)

/* Test data.
people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Status: "Prospect", Address: &Address{City: "City X", State: "State X"}})
people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
*/

func main() {
	router := mux.NewRouter()

	c := handlers.NewController()

	router.HandleFunc("/healthcheck", c.Healthcheck).Methods("GET")
	router.HandleFunc("/people", c.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", c.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}/status", c.GetStatus).Methods("GET")
	router.HandleFunc("/people/{id}/status/{newstatus}", c.SetStatus).Methods("PUT")
	router.HandleFunc("/people/{id}", c.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", c.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
