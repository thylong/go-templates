package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thylong/go-templates/05-gin-templ-htmx/web/views"
)

func SetupRoutes(r *gin.Engine) {
	// Healthcheck endpoint for k8s probes
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", views.Counter("testong"))
	})
}
