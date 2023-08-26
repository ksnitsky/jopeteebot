package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func requestToTgAPI(chatId string, message string, function string) (*http.Request, error) {
	accessToken := os.Getenv("TELEGRAM_BOT_ACCESS_TOKEN")
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/%s", accessToken, function)

	requestBody := map[string]interface{} {
		"chat_id": chatId,
		"text": message,
	}

	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{ "Content-Type": "application/json" }

	req, err := createHTTPRequest(apiURL, "POST", jsonBytes, headers)
	if err != nil {
		return nil, err
	}

	return req, nil
}