package handlers

import (
	"fmt"
)

func SendMessage(chatId int64, message string) error {
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

	return nil
}
