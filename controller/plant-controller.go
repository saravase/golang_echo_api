package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/saravase/golang_echo_api/entity"
	"github.com/saravase/golang_echo_api/service"
)

type PlantController interface {
	FindAll() []entity.Plant
	Save(ctx echo.Context) entity.Plant
}

type controller struct {
	service service.PlantService
}

func New(service service.PlantService) PlantController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Plant {
	return c.service.FindAll()
}

func (c *controller) Save(ctx echo.Context) entity.Plant {
	var plant entity.Plant
	ctx.Bind(&plant)
	c.service.Save(plant)
	return plant
}
