package main

import (
	"net/http"
	"smart-parking-lot/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var parks = []db.Park{
	{LotId: "A001", Licenseplate: "", Status: 0, Reserveable: false, Floor: "A"},
	{LotId: "A002", Licenseplate: "", Status: 1, Reserveable: false, Floor: "A"},
	{LotId: "A003", Licenseplate: "", Status: 0, Reserveable: false, Floor: "A"},
	{LotId: "A004", Licenseplate: "", Status: 0, Reserveable: false, Floor: "A"},
	{LotId: "A005", Licenseplate: "", Status: 0, Reserveable: false, Floor: "A"},
	{LotId: "A006", Licenseplate: "", Status: 0, Reserveable: false, Floor: "A"},
	{LotId: "A007", Licenseplate: "", Status: 0, Reserveable: false, Floor: "A"},
	{LotId: "A008", Licenseplate: "", Status: 0, Reserveable: true, Floor: "A"},
	{LotId: "A009", Licenseplate: "", Status: 0, Reserveable: true, Floor: "A"},
	{LotId: "A010", Licenseplate: "", Status: 0, Reserveable: true, Floor: "A"},
}

func main() {

	db.DB()
	// test webhook 4
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})

	e.GET("/parks", func(c echo.Context) error {

		var parks []db.Park

		db.DB().Find(&parks)

		return c.JSON(http.StatusOK, &parks)
	})

	e.POST("/create-mock-parks", func(c echo.Context) error {
		parks := parks

		db.DB().Create(&parks)

		return c.JSON(http.StatusCreated, parks)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
