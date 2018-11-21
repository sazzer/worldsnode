package internal

import (
	"time"

	"github.com/sirupsen/logrus"
	"grahamcox.co.uk/worlds/service/internal/health"
	"grahamcox.co.uk/worlds/service/internal/version"
	"grahamcox.co.uk/worlds/service/internal/service"
)

type dummyHealthcheck struct {
	result string
}

func (d dummyHealthcheck) CheckHealth(c chan<- health.Healthcheck) {
	c <- health.Healthcheck{
		Component:   "dummy",
		Measurement: "check",
		Status:      d.result,
	}
	time.Sleep(100 * time.Millisecond)
	c <- health.Healthcheck{
		Component:   "dummy",
		Measurement: "check",
		Status:      "PASS",
	}
	time.Sleep(300 * time.Millisecond)
	c <- health.Healthcheck{
		Component:   "dummy",
		Measurement: "check2",
		Status:      "PASS",
	}
}

// Main is the main entry point into the application
func Main(config Config) {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)

	healthchecker := health.New(
		dummyHealthcheck{health.Pass},
	)

	service := service.New(config.HTTP)
	service.AddRoutes(healthchecker.DefineRoutes)
	service.AddRoutes(version.DefineRoutes)
	service.Start()
}
