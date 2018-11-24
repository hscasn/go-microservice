package main

import (
	"go-microservice/internal/api"
	"go-microservice/internal/config"
	"go-microservice/internal/ioutils/client1"
	"go-microservice/internal/ioutils/client2"
	"go-microservice/pkg/health"
	"go-microservice/pkg/log"
	"go-microservice/pkg/server"
)

func main() {
	config := config.Create()
	log := log.Create(config.Name, false)

	onClose := func() {
		log.Infof("Server %s is shutting down\n", config.Name)
	}

	c1 := client1.Create()
	c2 := client2.Create()

	healthChecks := health.Checks{
		"client1": c1,
		"client2": c2,
	}

	srv := server.Create(log, healthChecks, 8000, onClose)
	api.Create(srv.Router)
	srv.Start()
}
