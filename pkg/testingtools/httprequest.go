package testingtools

import (
	"io/ioutil"
	"net/http"
)

// HTTPRequest can be used to send HTTP requests to a server. It will
// return the response and code
func HTTPRequest(
	url string,
	method string,
	path string,
) (*http.Response, string, error) {
	req, err := http.NewRequest(method, url+path, nil)
	if err != nil {
		return nil, "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()

	rbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}

	return res, string(rbody), nil
}
