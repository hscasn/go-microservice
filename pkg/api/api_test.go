package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"go-microservice/pkg/health"
	"go-microservice/pkg/testingtools"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	Create(router, health.Checks{})
	s := httptest.NewServer(router)
	defer s.Close()

	// Ready
	res, _ := testingtools.HTTPRequest(t, s.URL, "GET", "/ready")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}

	// Health
	res, _ = testingtools.HTTPRequest(t, s.URL, "GET", "/health")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	res, _ = testingtools.HTTPRequest(t, s.URL, "GET", "/health/details")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}

	// Settings
	res, _ = testingtools.HTTPRequest(t, s.URL, "GET", "/settings/loglevel")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	loglevels := []string{"", "debug", "info", "warn", "error", "fatal"}
	for _, p := range loglevels {
		path := fmt.Sprintf("/settings/loglevel/%s", p)
		res, _ = testingtools.HTTPRequest(t, s.URL, "PUT", path)
		if res.StatusCode != http.StatusOK {
			t.Errorf("should get OK status for level '%s'", p)
		}
	}
}
