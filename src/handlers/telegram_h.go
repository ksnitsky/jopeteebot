package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	tg "github.com/ksnitsky/jopeteebot/src/models/telegram_m"
	"github.com/labstack/echo/v4"
)

func SendMessageToTgAPI(chatId int64, message string) ([]byte, error) {
	req, err := requestToTgAPI(chatId, message, "sendMessage")
	if err != nil {
		return nil, err
	}

	resp, err := fetchAPI(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func requestToTgAPI(chatId int64, message string, function string) (*http.Request, error) {
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

func GetUpdates(c echo.Context) error {
	var newTgMessage tg.Update
	err := c.Bind(&newTgMessage)
	if err != nil {
		return err
	}

	err = SendMessage(newTgMessage.Message.From.ID, newTgMessage.Message.Text)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
