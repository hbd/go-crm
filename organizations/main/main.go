package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const ()

func main() {
	router := mux.NewRouter()

	c := organizations.NewController()

	router.HandleFunc("/healthcheck", c.Healthcheck).Methods(http.MethodGet)

	logrus.Fatal(http.ListenAndServe(":8000", nil))
}
