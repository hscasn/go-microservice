package health

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go-microservice/pkg/apiresponse"
	"go-microservice/pkg/health"
	"go-microservice/pkg/testingtools"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	Create(router, health.Checks{})
	s := httptest.NewServer(router)
	defer s.Close()

	res, _ := testingtools.HTTPRequest(t, s.URL, "GET", "/")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	res, _ = testingtools.HTTPRequest(t, s.URL, "GET", "/details")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
}

type dummyCheckThatSucceeds struct{}

func (d *dummyCheckThatSucceeds) Ping() bool {
	return true
}

type dummyCheckThatFails struct{}

func (d *dummyCheckThatFails) Ping() bool {
	return false
}

func TestSummarizedNoClients(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := summarizedController(health.Checks{})

	h(w, r)

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("should have OK status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.HEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
}

func TestSummarizedTwoThatSucceed(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := summarizedController(health.Checks{
		"one": &dummyCheckThatSucceeds{},
		"two": &dummyCheckThatSucceeds{},
	})

	h(w, r)

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("should have OK status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.HEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
}

func TestSummarizedOneThatFails(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := summarizedController(health.Checks{
		"one": &dummyCheckThatSucceeds{},
		"two": &dummyCheckThatFails{},
	})

	h(w, r)

	failedToPingRex, _ := regexp.Compile(
		"one or more health checks could not be pinged")

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 503 {
		t.Errorf("should have Service Unavailable status code header")
	}
	if res.Status.AdditionalInfo == nil {
		t.Errorf("response should include additional info")
	}
	for _, err := range res.Status.AdditionalInfo.Errors {
		foundErr := false
		if !failedToPingRex.MatchString(err) {
			foundErr = true
		}
		if !foundErr {
			t.Errorf(
				"response should an error explaining that a " +
					"health check could not be pinged")
		}
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.UNHEALTHY {
		t.Errorf(
			"response should have unhealthy status, but is \"%s\"",
			res.Result.(string))
	}
}

func TestSummarizedTwoThatFail(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := summarizedController(health.Checks{
		"one": &dummyCheckThatFails{},
		"two": &dummyCheckThatFails{},
	})

	h(w, r)

	failedToPingRex, _ := regexp.Compile(
		"one or more health checks could not be pinged")

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 503 {
		t.Errorf("should have Service Unavailable status code header")
	}
	if res.Status.AdditionalInfo == nil {
		t.Errorf("response should include additional info")
	}
	for _, err := range res.Status.AdditionalInfo.Errors {
		foundErr := false
		if !failedToPingRex.MatchString(err) {
			foundErr = true
		}
		if !foundErr {
			t.Errorf(
				"response should an error explaining that a " +
					"health check could not be pinged")
		}
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.UNHEALTHY {
		t.Errorf(
			"response should have unhealthy status, but is \"%s\"",
			res.Result.(string))
	}
}

func TestDetailsNoClients(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := detailsController(health.Checks{})

	h(w, r)

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("should have OK status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resResult := res.Result.(map[string]interface{})
	resHealth := resResult["status"].(string)
	if resHealth != health.HEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
	resSummary := resResult["healthSummary"].(map[string]interface{})
	if len(resSummary) > 0 {
		t.Errorf("expected health summary to be empty")
	}

}

func TestDetailsTwoThatSucceed(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := detailsController(health.Checks{
		"one": &dummyCheckThatSucceeds{},
		"two": &dummyCheckThatSucceeds{},
	})

	h(w, r)

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("should have OK status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resResult := res.Result.(map[string]interface{})
	resHealth := resResult["status"].(string)
	if resHealth != health.HEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
	resSummary := resResult["healthSummary"].(map[string]interface{})
	if len(resSummary) != 2 {
		t.Errorf("expected health summary to have two entries")
	}
	for srvName, srvHealth := range resSummary {
		if srvName == "one" && srvHealth == health.HEALTHY {
			continue
		} else if srvName == "two" && srvHealth == health.HEALTHY {
			continue
		} else {
			t.Errorf(
				"expected health summary to list "+
					"one->healthy, two->healthy"+
					"but got %s->%s",
				srvName,
				srvHealth)
		}
	}
}

func TestDetailsOneThatFails(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := detailsController(health.Checks{
		"one": &dummyCheckThatSucceeds{},
		"two": &dummyCheckThatFails{},
	})

	h(w, r)

	failedToPingRex, _ := regexp.Compile(
		"one or more health checks could not be pinged")

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 503 {
		t.Errorf("should have Service Unavailable status code header")
	}
	if res.Status.AdditionalInfo == nil {
		t.Errorf("response should include additional info")
	}
	for _, err := range res.Status.AdditionalInfo.Errors {
		foundErr := false
		if !failedToPingRex.MatchString(err) {
			foundErr = true
		}
		if !foundErr {
			t.Errorf(
				"response should an error explaining that a " +
					"health check could not be pinged")
		}
	}
	resResult := res.Result.(map[string]interface{})
	resHealth := resResult["status"].(string)
	if resHealth != health.UNHEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
	resSummary := resResult["healthSummary"].(map[string]interface{})
	if len(resSummary) != 2 {
		t.Errorf("expected health summary to have two entries")
	}
	for srvName, srvHealth := range resSummary {
		if srvName == "one" && srvHealth == health.HEALTHY {
			continue
		} else if srvName == "two" && srvHealth == health.UNHEALTHY {
			continue
		} else {
			t.Errorf(
				"expected health summary to list "+
					"one->healthy, two->unhealthy"+
					"but got %s->%s",
				srvName,
				srvHealth)
		}
	}
}

func TestDetailsTwoThatFail(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := detailsController(health.Checks{
		"one": &dummyCheckThatFails{},
		"two": &dummyCheckThatFails{},
	})

	h(w, r)

	failedToPingRex, _ := regexp.Compile(
		"one or more health checks could not be pinged")

	res := &apiresponse.APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 503 {
		t.Errorf("should have Service Unavailable status code header")
	}
	if res.Status.AdditionalInfo == nil {
		t.Errorf("response should include additional info")
	}
	for _, err := range res.Status.AdditionalInfo.Errors {
		foundErr := false
		if !failedToPingRex.MatchString(err) {
			foundErr = true
		}
		if !foundErr {
			t.Errorf(
				"response should an error explaining that a " +
					"health check could not be pinged")
		}
	}
	resResult := res.Result.(map[string]interface{})
	resHealth := resResult["status"].(string)
	if resHealth != health.UNHEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
	resSummary := resResult["healthSummary"].(map[string]interface{})
	if len(resSummary) != 2 {
		t.Errorf("expected health summary to have two entries")
	}
	for srvName, srvHealth := range resSummary {
		if srvName == "one" && srvHealth == health.UNHEALTHY {
			continue
		} else if srvName == "two" && srvHealth == health.UNHEALTHY {
			continue
		} else {
			t.Errorf(
				"expected health summary to list "+
					"one->unhealthy, two->unhealthy"+
					"but got %s->%s",
				srvName,
				srvHealth)
		}
	}
}
