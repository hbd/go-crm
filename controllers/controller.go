package controllers

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// Controller .
type Controller struct {
	db *sql.DB
}

// Credentials .
type Credentials struct {
	Host   string
	Port   int
	User   string
	Dbname string
}

// NewController .
func NewController() *Controller {
	var c Controller

	// Open a connection to the DB.
	creds := newCredentials()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		creds.Host, creds.Port, creds.User, creds.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	c.db = db

	return &c
}

// newCredentials returns new credentials.
// Pass in configurable credentials from the env?
func newCredentials() Credentials {
	return Credentials{
		Host:   "localhost",
		Port:   5432,
		User:   "postgres",
		Dbname: "crm",
	}
}

// PingDB pings the database.
func (c *Controller) PingDB() error {
	if err := c.db.Ping(); err != nil {
		return errors.Wrap(err, "db ping")
	}
	return nil
}
