package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func SendMessageToTgAPI(c *echo.Context) {
	// chatId := "153576749"

}

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

// func processTgResponse(resp *http.Response) (string, error) {
// 	contentType := resp.Header.Get("Content-Type")

// 	if strings.Contains(contentType, "application/json") {
// 		var responseData map[string]interface{}
// 		err := json.NewDecoder(resp.Body).Decode(&responseData)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return responseData, nil
// 	} else {
// 		responseBody, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return string(responseBody), nil
// 	}
// }
