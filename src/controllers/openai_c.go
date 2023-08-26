package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	m "github.com/ksnitsky/jopeteebot/src/models"
)

func requestToChatGPT(message string) (*http.Request, error) {
	apiKey := os.Getenv("OPENAI_API_TOKEN")
	apiURL := "https://api.openai.com/v1/chat/completions"

	requestBody := []byte(fmt.Sprintf(`{
		"model": "gpt-3.5-turbo",
		"messages": [
			{
				"role": "system",
				"content": "You are a helpful assistant."
			},
			{
				"role": "user",
				"content": "%s"
			}
		]
	}`, message))
	headers := map[string]string{
		"Content-Type": "application/json",
		"Authorization": "Bearer "+apiKey,
	}

	req, err := createHTTPRequest(apiURL, "POST", requestBody, headers)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func processResponse(resp *http.Response) (m.GPTResponse, error) {
	contentType := resp.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		var responseData map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	} else {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return string(responseBody), nil
	}
}

