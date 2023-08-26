package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

// TODO:
//  - longpoling to a /getUpdates to receive messages
//  - webhook
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLOWORLD")
	})

	e.GET("/sendMessage", sendMessage)

	e.Logger.Fatal(e.Start(":3000"))
}

// chat_id 153576749
// e.GET("/sendMessage", sendMessage)
func sendMessage(c echo.Context) error {
	// message := c.QueryParam("text")
	// chatId := c.QueryParam("chat_id")

	message := "Please name 10 Europe capitals"
	chatId := "153576749"

	ch := make(chan string)

	chatGPTreq, err := requestToChatGPT(message)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	go func() {
		openAiResponse, err := fetchAPI(chatGPTreq)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		openAiResponseData, err := controllers.processResponse(openAiResponse)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}


		tgReq, err := requestToTgAPI(chatId, string(openAiResponse), "sendMessage")
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}

	// body, err := processResponse(resp)

		tgResponse, err := fetchAPI(tgReq)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		ch <- string(tgResponse)
	}()

	return c.String(http.StatusOK, <- ch)
}

func fetchAPI(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
