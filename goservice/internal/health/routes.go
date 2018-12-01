package health

import "github.com/gin-gonic/gin"

// DefineRoutes will add the routes for the HealthChecks endpoints
func (h *Healthchecker) DefineRoutes(r *gin.RouterGroup) {
	r.GET("/health", h.get)
}
