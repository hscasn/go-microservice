package health

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/pkg/apiresponse"
	"github.com/hscasn/go-microservice/pkg/health"
)

// Create will bind this API to an existing router
func New(router chi.Router, healthChecks health.Checks) {
	router.Get("/", summarizedController(healthChecks))
	router.Get("/details", detailsController(healthChecks))
}

type simpleResponse struct {
	Status string `json:"status"`
}

type detailedResponse struct {
	Status        string            `json:"status"`
	HealthSummary map[string]string `json:"healthSummary"`
}

func summarizedController(healthChecks health.Checks) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendStatus(healthChecks, false, w)
	}
}

func detailsController(healthChecks health.Checks) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendStatus(healthChecks, true, w)
	}
}

// Common code for Default and Details controllers. Use the "detailed" boolean
// flag to switch from summarized to detailed
func sendStatus(
	healthChecks health.Checks,
	detailed bool,
	w http.ResponseWriter,
) {

	detailedHealth := health.New(healthChecks)
	summaryHealth := health.Summarize(detailedHealth)

	var s interface{} = simpleResponse{summaryHealth}
	if detailed {
		s = detailedResponse{summaryHealth, detailedHealth}
	}

	rData := apiresponse.ResponseData{
		Result: s,
	}

	if summaryHealth != "healthy" {
		rData.Code = http.StatusServiceUnavailable
		rData.Errors = append(
			rData.Errors,
			"One or more health checks could not be pinged")
	}

	apiresponse.SendJSONResponse(rData, w)
}
