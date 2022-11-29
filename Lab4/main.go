package main

import (
	"minimal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())

	controller := controller.NewWaterController()

	api := e.Group("/api")
	api.GET("/registers/:topic", controller.GetRegisters)
	api.POST("/subscribe/:topic", controller.SubscribeTopic)

	e.Logger.Fatal(e.Start(":1323"))
}
