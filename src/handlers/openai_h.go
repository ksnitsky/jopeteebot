package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	m "github.com/ksnitsky/jopeteebot/src/models"
)

func SendMessageToChatGPT(message string) (string, error) {
	req, err := requestToChatGPT(message)
	if err != nil {
		return "", err
	}
	resp, err := fetchAPI(req)
	if err != nil {
		return "", err
	}
	data, err := processGPTResponse(resp)
	if err != nil {
		return "", err
	}

	return data, nil
}

func requestToChatGPT(message string) (*http.Request, error) {
	apiKey := os.Getenv("OPENAI_API_TOKEN")
	apiURL := "https://api.openai.com/v1/chat/completions"
	escapedMessage := strings.ReplaceAll(message, `"`, `\"`)

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
	}`, escapedMessage))

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + apiKey,
	}

	req, err := createHTTPRequest(apiURL, "POST", requestBody, headers)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func processGPTResponse(body []byte) (string, error) {
	var gptResponse m.GPTResponse
	err := json.Unmarshal(body, &gptResponse)
	if err != nil {
		return "", err
	}

	message := gptResponse.Choices[0].Message.Content
	return message, nil
}
