package api

import (
	"github.com/go-chi/chi"
	"go-microservice/internal/api/dummy"
)

// Create will bind this API to an exiting router
func Create(router chi.Router) {
	router.Route("/dummy", func(r chi.Router) {
		dummy.Create(r)
	})
}
