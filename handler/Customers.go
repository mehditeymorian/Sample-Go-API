package handler

import (
	"GoServer/db"
	"GoServer/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateCustomer(c echo.Context) error {
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	inserted := db.Database.Insert(customer)
	return c.JSON(http.StatusOK, model.EchoResp(inserted, "success"))
}

func EditCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("cID"))
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}

	edited, err := db.Database.Edit(int64(id), customer)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.MsgResp("error (cID is not available)"))
	}
	return c.JSON(http.StatusOK, model.EchoResp(edited, "success"))
}

func DeleteCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("cID"))

	if err != nil {
		return err
	}

	err = db.Database.Delete(int64(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, model.MsgResp("error (cID is not available)"))
	}
	return c.JSON(http.StatusOK, model.MsgResp("success"))
}

func RetrieveAllCustomer(c echo.Context) error {
	customers, err := db.Database.RetrieveAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, model.MsgResp("error (customers are not available)"))
	}
	return c.JSON(http.StatusOK, model.RetrieveAllResp(customers, "success"))
}
