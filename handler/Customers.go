package handler

import (
	"GoServer/db"
	"GoServer/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func CreateCustomer(c echo.Context) error {
	// convert request body to customer
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	inserted := db.Database.Insert(customer)
	return c.JSON(http.StatusOK, model.EchoResp(inserted, "success"))
}

func EditCustomer(c echo.Context) error {
	// read id value from path and bind request body to customer
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
	// read id value from path
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

func RetrieveCustomer(c echo.Context) error {
	// read query parameter from path
	nameLike := c.QueryParam("cName")
	if len(nameLike) == 0 {
		return retrieveAllCustomer(c)
	}

	return retrieveSearchedCustomers(c, nameLike)
}

func retrieveAllCustomer(c echo.Context) error {
	customers, err := db.Database.RetrieveAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, model.MsgResp("error (customers are not available)"))
	}
	return c.JSON(http.StatusOK, model.RetrieveCustomersResp(customers, "success"))
}

// return first customer when customer.Name start with <nameLike>
func retrieveSearchedCustomers(c echo.Context, nameLike string) error {

	customers, err := db.Database.RetrieveAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, model.MsgResp("error (customers are not available)"))
	}

	var searchedCustomers []model.Customer
	for _, each := range customers {
		if strings.HasPrefix(each.Name, nameLike) {
			searchedCustomers = append(searchedCustomers, each)
		}
	}

	if searchedCustomers == nil {
		return c.JSON(http.StatusNotFound, "No Customer Found")
	}

	return c.JSON(http.StatusOK, model.RetrieveCustomersResp(searchedCustomers, "success"))
}
