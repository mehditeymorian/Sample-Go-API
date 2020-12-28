package handler

import (
	"GoServer/db"
	"GoServer/model"
	"GoServer/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateCustomer(c echo.Context) error {
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	inserted := db.Database.Insert(customer)
	return c.JSON(http.StatusOK, responses.EchoResp(inserted, "success"))
}

func EditCustomer(c echo.Context) error {
	return nil
}

func DeleteCustomer(c echo.Context) error {
	return nil
}

func RetrieveCustomer(c echo.Context) error {
	return c.JSON(http.StatusOK, responses.RetrieveAllResp(db.Database.RetrieveAll(), "success"))
}
