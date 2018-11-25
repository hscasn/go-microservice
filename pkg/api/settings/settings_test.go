package settings

import (
	"fmt"
	"github.com/go-chi/chi"
	"go-microservice/pkg/testingtools"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	Create(router)
	s := httptest.NewServer(router)
	defer s.Close()

	// Settings
	res, _, err := testingtools.HTTPRequest(s.URL, "GET", "/loglevel")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	loglevels := []string{"", "debug", "info", "warn", "error", "fatal"}
	for _, p := range loglevels {
		path := fmt.Sprintf("/loglevel/%s", p)
		res, _, err = testingtools.HTTPRequest(s.URL, "PUT", path)
		if err != nil {
			t.Error(err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("should get OK status for level '%s'", p)
		}
	}
}
