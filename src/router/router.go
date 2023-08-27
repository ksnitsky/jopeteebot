package router

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ksnitsky/jopeteebot/src/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		apiV1RouteGroup.POST("/getUpdates", handlers.GetUpdates)
	}

	return e
}

func StartApp(port string) {
	e := Initialize()

	e.Logger.Fatal(e.Start(port))
}
