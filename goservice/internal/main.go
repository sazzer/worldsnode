package internal

import (
	"github.com/benbjohnson/clock"
	"grahamcox.co.uk/worlds/service/internal/oauth2/accessToken"
	"github.com/sirupsen/logrus"
	"grahamcox.co.uk/worlds/service/internal/health"
	"grahamcox.co.uk/worlds/service/internal/version"
	"grahamcox.co.uk/worlds/service/internal/service"
	"grahamcox.co.uk/worlds/service/internal/database"
)

// Main is the main entry point into the application
func Main(config Config) {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)

	db, err := database.New(config.Database)
	if err != nil {
		panic("Failed to open database connection. Aborting.")
	}

	systemClock := clock.New()

	accessTokenSerializer := accesstoken.NewSerializer("", systemClock)

	healthchecker := health.New(
		db,
	)

	service := service.New(config.HTTP)
	service.AddMiddleware(accesstoken.NewMiddleware(accessTokenSerializer))

	service.AddRoutes(healthchecker.DefineRoutes)
	service.AddRoutes(version.DefineRoutes)

	service.Start()
}
