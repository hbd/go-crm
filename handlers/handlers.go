package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Person .
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Status    string   `json:"status,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address .
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Healthcheck .
func Healthcheck() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

// GetPeople .
func GetPeople(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	fmt.Println("Successfully connected!")

	json.NewEncoder(w).Encode(people)

}

// GetPerson .
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// GetStatus .
func GetStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item.Status)
			return
		}
	}
	json.NewEncoder(w).Encode("")
}

// SetStatus .
func SetStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			item.Status = params["newstatus"]
			json.NewEncoder(w).Encode(people)
			return
		}
	}
	json.NewEncoder(w).Encode("User with given id not found")
}

// CreatePerson .
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson .
func DeletePerson(w http.ResponseWriter, r *http.Request) {
}
