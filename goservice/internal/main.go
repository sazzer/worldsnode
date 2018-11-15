package internal

import (
	"net/http"

	"github.com/go-chi/chi"
	"grahamcox.co.uk/worlds/service/internal/service"
)

// Main is the main entry point into the application
func Main() {
	service := service.New()
	service.AddRoutes(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi"))
		})
	})
	service.Start()
}
