package main

import (
	"GoServer/db"
	"GoServer/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	// init database
	db.Database.Init()

	e := echo.New()

	e.POST("/customers", handler.CreateCustomer)
	e.PUT("PUT /customers/{cID}", handler.EditCustomer)
	e.GET("/customer", handler.RetrieveCustomer)
	e.DELETE("/customers/{cID}", handler.DeleteCustomer)

	e.Logger.Fatal(e.Start(":8080"))
}
