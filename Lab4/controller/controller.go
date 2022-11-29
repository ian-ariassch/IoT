package controller

import (
	"minimal/mqtt"
	"minimal/repository"
	"minimal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WaterController struct {
	service service.WaterService
}

func NewWaterController() WaterController {
	wRepo := repository.NewWaterRepository()
	manager := mqtt.NewMqttManager()
	return WaterController{
		service: service.NewWaterService(wRepo, *manager),
	}
}

func (wc WaterController) GetRegisters(c echo.Context) error {
	topic := c.Param("topic")
	registers, err := wc.service.GetRegisters(topic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, registers)
}

func (wc WaterController) SubscribeTopic(c echo.Context) error {
	topic := c.Param("topic")
	err := wc.service.Subscribe(topic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
