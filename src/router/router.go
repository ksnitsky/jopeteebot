package router

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ksnitsky/jopeteebot/src/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func Initialize() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	apiV1RouteGroup := e.Group("api/v1")
	{
		apiV1RouteGroup.GET("", func(c echo.Context) error {
			return c.String(http.StatusOK, "HELLOWORLD")
		})

		apiV1RouteGroup.POST("/sendMessage", handlers.SendMessage)
	}

	return e
}

func StartApp(port string) {
	e := Initialize()

	e.Logger.Fatal(e.Start(port))
}
