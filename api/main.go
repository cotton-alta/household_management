package main

import (
	"api-module/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
) 

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api", controllers.GetTable())
	e.POST("/api/items", controllers.CreateItem())
	e.GET("/api/items/:item", controllers.GetItem())
	e.DELETE("/api/items/:item", controllers.DeleteItem())

	e.Start(":3001")
}