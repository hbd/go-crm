package people

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/lib/pq" // PG Driver.
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

	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
	}

	c.aws = session.Must(session.NewSession())

	if useLocalDB := os.Getenv("LOCAL_DB"); useLocalDB == "true" {
		logrus.WithField("LOCAL_DB", useLocalDB).Debug("Using local DB.")
		c.db = MustNewLocalDBClient()
	} else {
		logrus.WithField("LOCAL_DB", useLocalDB).Debug("Using SQL DB.")
		c.db = MustNewSQLClient()
	}

	return &c, nil
}
