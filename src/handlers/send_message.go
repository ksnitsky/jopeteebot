package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SendMessage(c echo.Context) error {
	message := c.FormValue("text")
	chatId := c.FormValue("chat_id")
	// message := "Please name 10 Europe capitals"
	// chatId := "153576749"

	ch := make(chan string)

	go func() {
		openAiResponse, err := SendMessageToChatGPT(message)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}

		tgAPIResponse, err := SendMessageToTgAPI(chatId, openAiResponse)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		ch <- string(tgAPIResponse)
	}()

	return c.String(http.StatusOK, <- ch)
}
