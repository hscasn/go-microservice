package apiresponse

import (
	"encoding/json"
	"math"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateDefault(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()

	SendJSONResponse(ResponseData{}, w)

	res := &APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("standard response should have OK status code header")
	}
	if res.Status.Code != 200 || res.Status.Message != "OK" {
		t.Errorf(
			"an OK header should have status code of 200 " +
				"and message \"OK\" in body")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf(
			"standard response should have null additionalInfo")
	}
	if res.Result != "" {
		t.Errorf(
			"standard response should have empty string as result")
	}
	if !isValidDateTime(res.ServerTimestamp) {
		t.Errorf(
			"the timestamp should have a maximum difference of " +
				"5 min")
	}
}

type dummyResult struct {
	Prop1 string       `json:"prop1"`
	Prop2 int          `json:"prop2"`
	Prop3 bool         `json:"prop3"`
	Prop4 *dummyResult `json:"prop4"`
}

func TestCreateComplex(t *testing.T) {
	t.Parallel()

	result := dummyResult{
		Prop1: "prop 1",
		Prop2: 2,
		Prop3: true,
		Prop4: &dummyResult{
			Prop1: "other prop 1",
		},
	}

	w := httptest.NewRecorder()

	c := ResponseData{
		Errors: []string{
			"this is error 1",
			"this is error 2",
		},
		Warnings: []string{
			"this is warning 1",
			"this is warning 2",
		},
		Code: 404,
		Headers: map[string]string{
			"x-header": "my header",
		},
		Result: result,
	}
	SendJSONResponse(c, w)

	res := &APIResponse{}
	json.Unmarshal([]byte(w.Body.String()), res)

	expectedResult, _ := json.Marshal(result)
	actualResult, _ := json.Marshal(res.Result)

	if w.Code != 404 {
		t.Errorf("response should have 404 status header")
	}
	if res.Status.Code != 404 || res.Status.Message != "Not Found" {
		t.Errorf(
			"a 404 header should have status code of 404 " +
				"and message \"Not Found\" in body")
	}
	if res.Status.AdditionalInfo == nil ||
		len(res.Status.AdditionalInfo.Errors) != 2 ||
		len(res.Status.AdditionalInfo.Warnings) != 2 {
		t.Errorf(
			"response should have 2 errors and 2 warnings in " +
				"additionalInfo")
	}
	if res.Status.AdditionalInfo.Errors[0] != "this is error 1" ||
		res.Status.AdditionalInfo.Errors[1] != "this is error 2" ||
		res.Status.AdditionalInfo.Warnings[0] != "this is warning 1" ||
		res.Status.AdditionalInfo.Warnings[1] != "this is warning 2" {
		t.Errorf("expected additional info to have correct messages")
	}
	if string(actualResult) != string(expectedResult) {
		t.Errorf("standard response should have expected JSON")
	}
	if !isValidDateTime(res.ServerTimestamp) {
		t.Errorf(
			"the timestamp should have a maximum " +
				"difference of 5 min")
	}
}

func isValidDateTime(d string) bool {
	parsed, _ := time.Parse(time.RFC3339, d)
	since := time.Since(parsed)
	return math.Abs(since.Minutes()) < 5
}
