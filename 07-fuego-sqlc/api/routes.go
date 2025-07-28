package api

import (
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/thylong/go-templates/07-fuego-sqlc/pkg/db"
)

type HealthzResponse struct {
	Message string `json:"message"`
}

func SetupRoutes(r *fuego.Server, queries *db.Queries) {
	// Healthcheck endpoint for k8s probes
	fuego.Get(r, "/healthz",
		func(c fuego.ContextNoBody) (HealthzResponse, error) {
			c.SetHeader("Content-Type", "application/json")
			return HealthzResponse{Message: "ok"}, nil
		},
		option.Description("Healthcheck endpoint"),
		option.Tags("healthz", "healthcheck", "probe"),
		option.OperationID("healthz"),
	)

	// api := r.Group("/api")
	// router := api.Group("/auth")
	// router.POST("/register", handlers.NewAuthHandler(queries).SignUpUser)
}
