package main

import (
	"net/http"
	"smart-parking-lot/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.DB()
	// test
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})
	e.Logger.Fatal(e.Start(":1234"))
}
