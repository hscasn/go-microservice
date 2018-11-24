package settings

import (
	"github.com/go-chi/chi"
	"go-microservice/pkg/api/settings/loglevel"
)

// Create will bind this API to an exiting router
func Create(router chi.Router) {
	router.Route("/loglevel", func(r chi.Router) {
		loglevel.Create(r)
	})
}
