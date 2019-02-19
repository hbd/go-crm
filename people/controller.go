package people

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/lib/pq" // PG Driver.
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Controller .
type Controller struct {
	aws *session.Session
	db  DB
}

// NewController .
func NewController() (*Controller, error) {
	var c Controller

	c.aws = session.Must(session.NewSession())

	if useLocalDB := os.Getenv("LOCAL_DB"); useLocalDB == "true" {

	} else {
		db, err := NewSQLClient()
		if err != nil {
			logrus.WithError(err).Debug("Error initializing a SQL client.")
			return &c, errors.Wrap(err, "new sql client")
		}
		c.db = db
	}

	return &c, nil
}
