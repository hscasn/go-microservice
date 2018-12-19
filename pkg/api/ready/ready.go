package ready

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/pkg/apiresponse"
	"github.com/hscasn/go-microservice/pkg/health"
)

// Create will bind this API to an existing router
func New(router chi.Router, healthChecks health.Checks) {
	router.Get("/", controller(healthChecks))
}

type readyResponse struct {
	Status string `json:"status"`
}

func controller(healthChecks health.Checks) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendStatus(healthChecks, w)
	}
}

func sendStatus(healthChecks health.Checks, w http.ResponseWriter) {
	status := health.CreateSummarized(healthChecks)

	s := readyResponse{status}

	rData := apiresponse.ResponseData{
		Result: s,
	}

	if status != "healthy" {
		rData.Code = http.StatusServiceUnavailable
	}

	apiresponse.SendJSONResponse(rData, w)
}
