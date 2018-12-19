package dummy

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/pkg/apiresponse"
)

// Create will bind this API to an exiting router
func New(router chi.Router) {
	router.Get("/", controller)
}

func controller(w http.ResponseWriter, r *http.Request) {
	apiresponse.SendJSONResponse(apiresponse.ResponseData{
		Result: "Hi! This is the dummy endpoint!",
	}, w)
}
