package version

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

var renderer *render.Render

func init() {
	renderer = render.New()
}

type versionResponse struct {
	Version string `json:"version"`
}

// DefineRoutes will add the routes for the Versions endpoints
func DefineRoutes(r chi.Router) {
	r.Get("/api/version", func(w http.ResponseWriter, r *http.Request) {
		renderer.JSON(w, 200, versionResponse{"1.0.0"})
	})
}
