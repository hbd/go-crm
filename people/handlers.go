package people

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var people []Person

// Healthcheck .
func (c *Controller) Healthcheck(w http.ResponseWriter, req *http.Request) {
	if err := c.PingDB(); err != nil {
		// Return error.
		logrus.WithError(err).Debugf("Failed to ping DB.")
		w.WriteHeader(http.StatusTeapot)
	}
	// Success.
	w.WriteHeader(http.StatusNoContent)
}

// GetPeople .
func (c *Controller) GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&people)
}

// GetPerson .
func (c *Controller) GetPerson(w http.ResponseWriter, r *http.Request) {
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
func (c *Controller) GetStatus(w http.ResponseWriter, r *http.Request) {
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
func (c *Controller) SetStatus(w http.ResponseWriter, r *http.Request) {
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
func (c *Controller) CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson .
func (c *Controller) DeletePerson(w http.ResponseWriter, r *http.Request) {
}
