package loglevel

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"go-microservice/pkg/apiresponse"
	"go-microservice/pkg/log"
	"go-microservice/pkg/testingtools"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	Create(router)
	s := httptest.NewServer(router)
	defer s.Close()

	// Settings
	res, _ := testingtools.HTTPRequest(t, s.URL, "GET", "/")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	loglevels := []string{"", "debug", "info", "warn", "error", "fatal"}
	for _, p := range loglevels {
		res, _ = testingtools.HTTPRequest(
			t,
			s.URL,
			"PUT",
			fmt.Sprintf("/%s", p))
		if res.StatusCode != http.StatusOK {
			t.Errorf("should get OK status for level '%s'", p)
		}
	}
}

func TestGetLevel(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := getLevel

	h(w, r)

	currLevelRex, _ := regexp.Compile(
		"(?i)current level: (debug|info|warn|error|fatal)")

	usageRex, _ := regexp.Compile(
		"(?i)specify a level by hitting the endpoint")

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("should have OK status code header")
	}
	if res.Status.AdditionalInfo == nil ||
		!usageRex.MatchString(res.Status.AdditionalInfo.Warnings[0]) {
		t.Errorf("response should include additional info for usage")
	}
	if !currLevelRex.MatchString(res.Result.(string)) {
		t.Errorf(
			"response should have current log level, but is \"%s\"",
			res.Result.(string))
	}
}

func TestPutLevelFor(t *testing.T) {
	t.Parallel()

	allLevels := []log.Level{
		log.DebugLevel,
		log.InfoLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.FatalLevel,
	}

	fallbackRex, _ := regexp.Compile(
		"(?i)level not specified; falling back")

	testLevel := func(level log.Level, def bool) {
		r := httptest.NewRequest("PUT", "http://localhost", nil)
		w := httptest.NewRecorder()
		h := putLevelFor(level, def)
		h(w, r)

		currLevelRex, _ := regexp.Compile(fmt.Sprintf(
			"(?i)level changed to: %s", level.String()))

		res := &apiresponse.APIResponse{}
		json.Unmarshal([]byte(w.Body.String()), res)

		testDefAdditionalInfo := func() {
			addInfo := res.Status.AdditionalInfo
			if addInfo == nil ||
				!fallbackRex.MatchString(addInfo.Warnings[0]) {
				t.Errorf(
					"response should include additional " +
						"info for usage")
			}
		}

		testNotDefAdditionalInfo := func() {
			if res.Status.AdditionalInfo != nil {
				t.Errorf(
					"response should not include " +
						"additional info")
			}
		}

		if w.Code != 200 {
			t.Errorf("should have OK status code header")
		}
		if def {
			testDefAdditionalInfo()
		} else {
			testNotDefAdditionalInfo()
		}
		if !currLevelRex.MatchString(res.Result.(string)) {
			t.Errorf(
				"response should log level %s, but "+
					"is \"%s\"",
				level.String(),
				res.Result.(string))
		}
	}

	for _, level := range allLevels {
		for _, def := range []bool{true, false} {
			testLevel(level, def)
		}
	}
}
