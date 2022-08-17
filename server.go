package main

import (
	"net/http"
	"smart-parking-lot/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.DB()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})
//test
	e.Logger.Fatal(e.Start(":1234"))
}
