package main

import (
       "log"
       "net/http"
       "github.com/gorilla/mux"

       "github.com/benjamin28/go-crm/handlers"
)

func main() {
     router := mux.NewRouter()

/*     people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Status: "Prospect", Address: &Address{City: "City X", State: "State X"}})
     people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
     people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})*/



     router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
     router.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
     router.HandleFunc("/people/{id}/status", handlers.GetStatus).Methods("GET")
     router.HandleFunc("/people/{id}/status/{newstatus}", handlers.SetStatus).Methods("PUT")

     router.HandleFunc("/people/{id}", handlers.CreatePerson).Methods("POST")
     router.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")

     log.Fatal(http.ListenAndServe(":8000", router))
}
