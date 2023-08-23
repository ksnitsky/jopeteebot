package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

// TODO:
//  - longpoling to a /getUpdates to receive messages
func main() {
	e := echo.New()

	// accessToken := os.Getenv("TELEGRAM_BOT_ACCESS_TOKEN")
	// requestUri := fmt.Sprintf("https://api.telegram.org/bot%s/%s", accessToken, "sendMessage")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLOWORLD")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
