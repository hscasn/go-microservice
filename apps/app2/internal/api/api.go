package api

import (
	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/apps/app2/internal/api/dummy"
)

// Create will bind this API to an exiting router
func Create(router chi.Router) {
	router.Route("/dummy", func(r chi.Router) {
		dummy.Create(r)
	})
}
