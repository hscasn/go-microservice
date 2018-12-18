package main

import (
	"github.com/hscasn/go-microservice/apps/app1/internal/api"
	"github.com/hscasn/go-microservice/apps/app1/internal/config"
	"github.com/hscasn/go-microservice/apps/app1/internal/dummyservice"
	"github.com/hscasn/go-microservice/pkg/health"
	"github.com/hscasn/go-microservice/pkg/log"
	"github.com/hscasn/go-microservice/pkg/server"
)

func main() {
	config := config.Create()
	log := log.Create(config.Name, false)

	onClose := func() {
		log.Infof("Server %s is shutting down\n", config.Name)
	}

	s := dummyservice.New()

	healthChecks := health.Checks{
		"dummyworker1": s,
	}

	srv := server.Create(log, healthChecks, 8000, onClose)
	api.Create(srv.Router)
	srv.Start()
}
