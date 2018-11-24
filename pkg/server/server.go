package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"go-microservice/pkg/api"
	"go-microservice/pkg/health"
	"go-microservice/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server contains the settings for the server
type Server struct {
	log      log.Interface
	Router   *chi.Mux
	onClose  func()
	shutdown chan bool
	httpSrv  *http.Server
}

// Create a new server based on a list of services
func Create(
	log log.Interface,
	healthChecks health.Checks,
	onClose func(),
) *Server {
	router := chi.NewRouter()
	server := &Server{log, router, onClose, nil, nil}

	api.Create(router, healthChecks)

	return server
}

// Start a server
func (s *Server) Start() {
	shutdown := make(chan bool, 1)
	s.shutdown = shutdown

	// Capturing OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		s.log.Warnf("Received OS signal %s. Shutting down", sig)
		s.onClose()
		shutdown <- true
	}()

	go func() {

		addr := fmt.Sprintf("0.0.0.0:%d", 8000)
		srv := &http.Server{
			Addr:         addr,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      s.Router,
		}
		s.httpSrv = srv
		s.log.Infof("Server starting at %s", addr)
		if err := srv.ListenAndServe(); err != nil {
			s.onClose()
			s.log.Error(err)
			s.log.Warn("Server received an error. Shutting down")
			shutdown <- true
		}
	}()

	<-shutdown
	s.onClose()
}
