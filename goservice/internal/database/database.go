package database

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	// We need to import the Postgres drivers but we don't ever use it directly
	_ "github.com/lib/pq"
)

// Database is the structure representing the actual database wrapper for our application
type Database struct {
	client *sql.DB
}

// New will create a new Database wrapper for our application
func New(config Config) (*Database, error) {
	db, err := sql.Open("postgres", config.URL)

	if err != nil {
		logrus.
			WithField("err", err).
			WithField("url", config.URL).
			Warn("Failed to open database connection")
		return nil, err
	}

	logrus.
		WithField("url", config.URL).
		Info("Opened database connection")
	return &Database{db}, nil
}
