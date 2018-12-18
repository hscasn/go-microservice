package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hscasn/go-microservice/pkg/health"

	"github.com/hscasn/go-microservice/pkg/api"
	"github.com/hscasn/go-microservice/pkg/log"

	"github.com/go-chi/chi"
)

// Server contains the settings for the server
type Server struct {
	log      log.Interface
	Router   *chi.Mux
	onClose  func()
	Shutdown chan bool
	httpSrv  *http.Server
	Addr     string
	port     int
}

// Create creates a new server based on a list of services
func Create(
	log log.Interface,
	healthChecks health.Checks,
	port int,
	onClose func(),
) *Server {
	router := chi.NewRouter()
	server := &Server{
		log:     log,
		Router:  router,
		onClose: onClose,
		port:    port,
	}

	api.Create(router, healthChecks)

	return server
}

// Start a server
func (s *Server) Start() {
	shutdown := make(chan bool, 1)
	s.Shutdown = shutdown

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
		s.Addr = fmt.Sprintf("0.0.0.0:%d", s.port)
		srv := &http.Server{
			Addr:         s.Addr,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      s.Router,
		}
		s.httpSrv = srv
		s.log.Infof("Server starting at %s", s.Addr)
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
