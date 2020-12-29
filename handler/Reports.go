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
