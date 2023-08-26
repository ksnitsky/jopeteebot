package controllers

import (
	"bytes"
	"net/http"
)

func createHTTPRequest(url, method string, body []byte, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}
