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
		return c.String(http.StatusOK, "Hello")
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

	e.PUT("/parks/:id", func(c echo.Context) error {
		park := db.Park{}
		if err := c.Bind(&park); err != nil {
			return err
		}

		parkDB := db.Park{
			ID:           park.ID,
			LotId:        park.LotId,
			Licenseplate: park.Licenseplate,
			Status:       park.Status,
			Reserveable:  park.Reserveable,
			Floor:        park.Floor,
		}

		id := c.Param("id")

		if err := db.DB().Model(&parkDB).Where("id = ?", id).Updates(&parkDB).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &parkDB)
	})

	e.GET("/deleteParks/:id", func(c echo.Context) error {
		park := db.Park{}
		id := c.Param("id")

		if err := db.DB().Delete(&park, id).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, id)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
