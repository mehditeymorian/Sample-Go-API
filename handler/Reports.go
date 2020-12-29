package handler

import (
	"GoServer/db"
	"GoServer/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func MonthlyReport(c echo.Context) error {
	month, _ := strconv.Atoi(c.Param("month"))

	if month > 11 {
		return c.JSON(http.StatusBadRequest, "Months are zero base. range from 0 to 11")
	}

	customers, err := db.Database.RetrieveAll()

	if err != nil {
		return c.JSON(http.StatusNotFound, model.MsgResp("no customer found"))
	}

	count := 0
	for _, each := range customers {
		eachMonth, _ := strconv.Atoi(strings.Split(each.RegisterDate, "-")[1])
		if eachMonth == month+1 {
			count++
		}
	}

	result := model.MonthlyReport{
		TotalCustomers: int64(count),
		Period:         int64(month),
		Msg:            "success",
	}

	return c.JSON(http.StatusOK, result)
}

func AnnualReport(c echo.Context) error {
	customers, err := db.Database.RetrieveAll()

	if err != nil {
		return c.JSON(http.StatusNotFound, model.MsgResp("no customer found"))
	}

	var count [12]int64
	for _, each := range customers {
		eachMonth, _ := strconv.Atoi(strings.Split(each.RegisterDate, "-")[1])
		count[eachMonth-1]++
	}

	var months [12]model.MonthInAnnualReport
	for i, eachMonthCustomers := range count {
		months[i] = model.MonthInAnnualReport{
			TotalCustomers: eachMonthCustomers,
			Period:         int64(i),
		}
	}

	result := model.AnnualReport{
		Months: months,
		Msg:    "success",
	}

	return c.JSON(http.StatusOK, result)
}
