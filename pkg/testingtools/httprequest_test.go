package testingtools

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	router.Patch("/my/path", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(fmt.Sprintf("hello world")))
	})
	s := httptest.NewServer(router)
	defer s.Close()

	res, body := HTTPRequest(t, s.URL, "PATCH", "/my/path")

	if res.StatusCode != http.StatusTeapot {
		t.Errorf("status code should be 'i am a teapot'")
	}
	if body != "hello world" {
		t.Errorf("body should be 'hello world'")
	}
}
