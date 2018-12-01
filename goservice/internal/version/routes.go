package version

import (
	"github.com/gin-gonic/gin"
)

type versionResponse struct {
	Version string `json:"version"`
}

// DefineRoutes will add the routes for the Versions endpoints
func DefineRoutes(r *gin.RouterGroup) {
	r.GET("/api/version", func(c *gin.Context) {
		c.JSON(200, versionResponse{"1.0.0"})
	})
}
