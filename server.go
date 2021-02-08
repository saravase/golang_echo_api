package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saravase/golang_echo_api/controller"
	"github.com/saravase/golang_echo_api/service"
)

var (
	plantService    service.PlantService       = service.New()
	plantController controller.PlantController = controller.New(plantService)
)

func main() {

	e := echo.New()

	e.GET("/plants", func(c echo.Context) error {
		return c.JSON(http.StatusOK, plantController.FindAll())
	})

	e.POST("/plants", func(c echo.Context) error {
		return c.JSON(http.StatusOK, plantController.Save(c))
	})

	e.Logger.Fatal(e.Start(":9090"))

}
