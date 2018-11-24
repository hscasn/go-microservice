package api

import (
	"github.com/go-chi/chi"
	"go-microservice/pkg/api/health"
	"go-microservice/pkg/api/ready"
	"go-microservice/pkg/api/settings"
	healthPkg "go-microservice/pkg/health"
)

// Create will bind this API to an existing router
func Create(router *chi.Mux, healthChecks healthPkg.Checks) {
	router.Route("/health", func(r chi.Router) {
		health.Create(r, healthChecks)
	})
	router.Route("/ready", func(r chi.Router) {
		ready.Create(r, healthChecks)
	})
	router.Route("/settings", func(r chi.Router) {
		settings.Create(r)
	})
}
