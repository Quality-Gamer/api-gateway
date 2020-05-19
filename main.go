package main

import (
	"gateway/request"
	"github.com/labstack/echo"
)

//main contains all API endpoints
func main() {
	e := echo.New()

	e.POST("/gateway/route", request.Route)

	//os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + "8005"))
}
