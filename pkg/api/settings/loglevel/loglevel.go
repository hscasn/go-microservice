package loglevel

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/pkg/apiresponse"
	"github.com/hscasn/go-microservice/pkg/log"
	"net/http"
)

// Create will bind this API to an exiting router
func Create(router chi.Router) {
	router.Get("/", getLevel)

	router.Put("/", putLevelFor(log.WarnLevel, true))
	router.Put("/debug", putLevelFor(log.DebugLevel, false))
	router.Put("/info", putLevelFor(log.InfoLevel, false))
	router.Put("/warn", putLevelFor(log.WarnLevel, false))
	router.Put("/error", putLevelFor(log.ErrorLevel, false))
	router.Put("/fatal", putLevelFor(log.FatalLevel, false))
}

type readyResponse struct {
	Status string `json:"status"`
}

var usage = "Specify a level by hitting the endpoint with /debug, /info, " +
	"/warn, /error, or /fatal with PUT method"

func getLevel(w http.ResponseWriter, r *http.Request) {
	level := log.GetLevel()
	message := fmt.Sprintf("Current level: %s", level.String())

	apiresponse.SendJSONResponse(apiresponse.ResponseData{
		Warnings: []string{usage},
		Result:   message,
	}, w)
}

func putLevelFor(level log.Level, isDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.SetLevel(level)

		message := fmt.Sprintf("Level changed to: %s", level.String())
		lvlNotSpecifiedMsg := fmt.Sprintf(
			"Level not specified; falling back to default. %s",
			usage)

		rData := apiresponse.ResponseData{
			Warnings: []string{},
			Result:   message,
		}

		if isDefault {
			rData.Warnings = append(
				rData.Warnings,
				lvlNotSpecifiedMsg)
		}

		apiresponse.SendJSONResponse(rData, w)
	}
}
