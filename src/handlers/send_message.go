package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// TODO:
//  - longpoling to a /getUpdates to receive messages
//  - webhook
// // e.POST("/sendMessage", sendMessage)
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

		// // openAiResponseData, err := controllers.processResponse(openAiResponse)
		// if err != nil {
		// 	ch <- fmt.Sprint(err)
		// 	return
		// }

		// tgReq, err := requestToTgAPI(chatId, string(openAiResponse), "sendMessage")
		// if err != nil {
		// 	ch <- fmt.Sprint(err)
		// 	return
		// }

	// body, err := processResponse(resp)

		// tgResponse, err := fetchAPI(tgReq)
		// if err != nil {
		// 	ch <- fmt.Sprint(err)
		// 	return
		// }
		// ch <- string(tgResponse)
