package people

import (
	"context"
)

var people []Person

// LocalDBClient .
type LocalDBClient struct{}

// NewLocalDBClient .
func NewLocalDBClient() (*LocalDBClient, error) {
	return &LocalDBClient{}, nil
}

// MustNewLocalDBClient .
func MustNewLocalDBClient() *LocalDBClient {
	return &LocalDBClient{}
}

/*
  General DB.
*/

// PingDB pings the database.
func (c *LocalDBClient) PingDB() error {
	return nil
}

/*
  Person.
*/

// GetPerson retrieves the person with the given ID from the DB.
func (c *LocalDBClient) GetPerson(ctx context.Context, id string) (Person, error) {
	for _, person := range people {
		if person.ID == id {
			return person, nil
		}
	}
	return Person{}, NewErrNotFound(id)
}

// InsertPerson adds a new person to the DB without checking if they exist first.
func (c *LocalDBClient) InsertPerson(ctx context.Context, person Person) error {
	people = append(people, person)
	return nil
}

// UpsertPerson inserts a person if they exist or updates their existing record.
func (c *LocalDBClient) UpsertPerson(ctx context.Context, person Person) error {
	return nil
}

// DeletePerson deletes the person with the given ID from the DB.
func (c *LocalDBClient) DeletePerson(ctx context.Context, id string) error {
	return nil
}

/*
  People.
*/

// GetPeople returns everyone in the DB.
func (c *LocalDBClient) GetPeople(ctx context.Context) ([]Person, error) {
	return people, nil
}
