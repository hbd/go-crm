package people

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Path parameters.
const (
	PathParamID        = "id"
	PathParamNewStatus = "newstatus"
)

/*
  People.
*/

// GetPeople .
func (c *Controller) GetPeople(w http.ResponseWriter, req *http.Request) {
	people, err := c.db.GetPeople(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)
}

/*
  Person.
*/

// CreatePerson .
func (c *Controller) CreatePerson(w http.ResponseWriter, req *http.Request) {
	var person Person
	if err := json.NewDecoder(req.Body).Decode(&person); err != nil {
		logrus.WithError(err).Debug("Error decoding person.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := getID(w, req)
	if err != nil {
		return
	}
	person.ID = id
	if err := c.db.InsertPerson(req.Context(), person); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetPerson .
func (c *Controller) GetPerson(w http.ResponseWriter, req *http.Request) {
	id, err := getID(w, req)
	if err != nil {
		return
	}
	person, err := c.db.GetPerson(req.Context(), id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Debugf("Error when getting person.")
		switch e := err.(type) {
		case *ErrNotFound:
			// TODO: Write header with middleware, e.g. use HTTP Error type and write its code to the response or 500 default.
			w.WriteHeader(e.Code())
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(person)
}

// DeletePerson .
func (c *Controller) DeletePerson(w http.ResponseWriter, r *http.Request) {
}

/*
  Status.
*/

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

/*
  General API.
*/

// Healthcheck .
func (c *Controller) Healthcheck(w http.ResponseWriter, req *http.Request) {
	if err := c.db.PingDB(); err != nil {
		// Return error.
		logrus.WithError(err).Debugf("Failed to ping DB.")
		w.WriteHeader(http.StatusTeapot)
	}
	// Success.
	w.WriteHeader(http.StatusNoContent)
}

func getID(w http.ResponseWriter, req *http.Request) (string, error) {
	params := mux.Vars(req)
	id, ok := params[PathParamID]
	if !ok {
		logrus.Debug("Invalid path param given.")
		w.WriteHeader(http.StatusBadRequest)
		return "", errors.New("id not found in path")
	}
	logrus.WithField("id", id).Debug("Retrieved ID.")
	return id, nil
}
