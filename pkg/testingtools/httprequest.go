package testingtools

import (
	"io/ioutil"
	"net/http"
	"testing"
)

// HTTPRequest can be used to send HTTP requests to a server. It will
// return the response and code
func HTTPRequest(
	t *testing.T,
	url string,
	method string,
	path string,
) (*http.Response, string) {

	req, err := http.NewRequest(method, url+path, nil)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	rbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	return res, string(rbody)
}
