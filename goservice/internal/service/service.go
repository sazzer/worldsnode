package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"

	log "github.com/sirupsen/logrus"
)

// Service is the actual Web Service
type Service struct {
	config Config
	router *gin.Engine
}

// New will create a new Web Service to work with
func New(config Config) Service {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(ginlogrus.Logger(log.New()))
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return Service{
		config,
		r,
	}
}

// AddRoutes will define routes on the service provided by the given Router Definition function
func (s *Service) AddRoutes(routeDefiner func(*gin.RouterGroup)) {
	routeDefiner(&s.router.RouterGroup)
}

// AddMiddleware will add the provided middleware to the service
func (s *Service) AddMiddleware(middleware func(*gin.Context)) {
	s.router.Use(middleware)
}

// Start will start the Web Service listening
func (s *Service) Start() {
	log.Info("Starting service")

	listenAddress := fmt.Sprintf(":%d", s.config.Port)
	log.WithField("address", listenAddress).Info("Service started")

	s.router.Run(listenAddress)
}
