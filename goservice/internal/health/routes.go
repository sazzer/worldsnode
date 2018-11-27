package health

import (
	"github.com/go-chi/chi"
)

// DefineRoutes will add the routes for the HealthChecks endpoints
func (h *Healthchecker) DefineRoutes(r chi.Router) {
	r.Get("/health", h.get)
}
