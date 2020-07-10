package main

import (
	"github.com/Quality-Gamer/api-gateway/auth"
	"github.com/Quality-Gamer/api-gateway/request"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

//main contains all API endpoints
func main() {
	e := echo.New()
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return auth.KeyValidator(key), nil
		},
	}))

	e.POST("/gateway/route", request.Route)

	//os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
