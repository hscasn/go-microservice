package api

import (
	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/apps/app2/internal/api/dummy"
)

// New will bind this API to an exiting router
func New(router chi.Router) {
	router.Route("/dummy", func(r chi.Router) {
		dummy.New(r)
	})
}
