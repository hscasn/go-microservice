package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/pkg/health"
	"github.com/hscasn/go-microservice/pkg/testingtools"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	New(router, health.Checks{})
	s := httptest.NewServer(router)
	defer s.Close()

	// Ready
	res, _, err := testingtools.HTTPRequest(s.URL, "GET", "/ready")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}

	// Health
	res, _, err = testingtools.HTTPRequest(s.URL, "GET", "/health")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	res, _, err = testingtools.HTTPRequest(s.URL, "GET", "/health/details")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}

	// Settings
	res, _, err = testingtools.HTTPRequest(s.URL, "GET", "/settings/loglevel")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	loglevels := []string{"", "debug", "info", "warn", "error", "fatal"}
	for _, p := range loglevels {
		path := fmt.Sprintf("/settings/loglevel/%s", p)
		res, _, err = testingtools.HTTPRequest(s.URL, "PUT", path)
		if err != nil {
			t.Error(err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("should get OK status for level '%s'", p)
		}
	}
}
