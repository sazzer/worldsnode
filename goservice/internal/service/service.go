package service

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	log "github.com/sirupsen/logrus"
)

// Service is the actual Web Service
type Service struct {
	router *chi.Mux
}

// New will create a new Web Service to work with
func New() Service {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.GetHead)
	router.Use(cors.Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	return Service{
		router,
	}
}

// AddRoutes will define routes on the service provided by the given Router Definition function
func (s *Service) AddRoutes(routeDefiner func(chi.Router)) {
	routeDefiner(s.router)
}

// AddMiddleware will add the provided middleware to the service
func (s *Service) AddMiddleware() {

}

// Start will start the Web Service listening
func (s *Service) Start() {
	log.Info("Starting service")

	http.ListenAndServe(":3000", s.router)
}
