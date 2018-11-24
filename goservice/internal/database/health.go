package database

import (
	"github.com/sirupsen/logrus"
	"grahamcox.co.uk/worlds/service/internal/health"
)

// CheckHealth will check that we can connect to the database and perform a trivial query
func (d *Database) CheckHealth(c chan<- health.Healthcheck) {
	err := d.client.Ping()
	if err == nil {
		c <- health.Healthcheck{
			Component:   "database",
			Measurement: "ping",
			Status:      health.Pass,
		}

		stats := d.client.Stats()
		c <- health.Healthcheck{
			Component:   "database",
			Measurement: "connections",
			Status:      health.Pass,
			Unit:        "open",
			Value:       stats.OpenConnections,
		}
		c <- health.Healthcheck{
			Component:   "database",
			Measurement: "connections",
			Status:      health.Pass,
			Unit:        "idle",
			Value:       stats.Idle,
		}
		c <- health.Healthcheck{
			Component:   "database",
			Measurement: "connections",
			Status:      health.Pass,
			Unit:        "inUse",
			Value:       stats.InUse,
		}
		c <- health.Healthcheck{
			Component:   "database",
			Measurement: "connections",
			Status:      health.Pass,
			Unit:        "max",
			Value:       stats.MaxOpenConnections,
		}
	} else {
		logrus.WithField("err", err).Warn("Database Ping failed")
		c <- health.Healthcheck{
			Component:   "database",
			Measurement: "ping",
			Status:      health.Fail,
		}
	}
}
