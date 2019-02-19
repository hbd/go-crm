package people

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Debug("Error initializing a SQL client.")
		return nil, errors.Wrap(err, "new sql client")
	}
	return &SQLClient{db}, nil
}

// MustNewSQLClient panics if initializing a new SQL Client fails.
func MustNewSQLClient() *SQLClient {
	db, err := NewSQLClient()
	if err != nil {
		logrus.WithError(err).Fatal("Error initializing new SQL Client.")
		return nil
	}
	return db
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
	return []Person{}, nil
}
