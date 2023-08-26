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
	// message := c.QueryParam("text")

	ch := make(chan string)

	// chatGPTreq, err := requestToChatGPT(message)
	// if err != nil {
	// 	ch <- fmt.Sprint(err)
	// }
	go func() {
		openAiResponse, err := SendMessageToChatGPT(c)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		ch <- openAiResponse
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