package people

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// SQLClient .
type SQLClient struct {
	*sql.DB
}

// Credentials .
type Credentials struct {
	Host   string
	Port   int
	User   string
	Dbname string
}

// NewSQLClient .
func NewSQLClient() (*SQLClient, error) {
	// Open a connection to the DB.
	creds := newCredentials()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		creds.Host, creds.Port, creds.User, creds.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return &SQLClient{db}, nil
}

// newCredentials returns new credentials.
// TODO: Make these dynamic by pulling from env using envconfig.
func newCredentials() Credentials {
	return Credentials{
		Host:   "localhost",
		Port:   5432,
		User:   "postgres",
		Dbname: "crm",
	}
}

// PingDB pings the database.
func (c *SQLClient) PingDB() error {
	if err := c.Ping(); err != nil {
		return errors.Wrap(err, "db ping")
	}
	return nil
}

// Queries.
const (
	queryGetPeople = `
SELECT * FROM people;
`
)

// GetPerson retrieves the person with the given ID from the DB.
func (c *SQLClient) GetPerson(ctx context.Context, id string) (Person, error) {
	return Person{}, nil
}

// UpsertPerson deletes the person with the given ID from the DB.
func (c *SQLClient) UpsertPerson(ctx context.Context, person Person) error {
	return nil
}

// DeletePerson deletes the person with the given ID from the DB.
func (c *SQLClient) DeletePerson(ctx context.Context, id string) error {
	return nil
}

// GetPeople returns everyone in the DB.
func (c *SQLClient) GetPeople(ctx context.Context) ([]Person, error) {
	return people, nil
}
