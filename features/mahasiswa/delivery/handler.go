package delivery

import (
	"be13/ca/features/mahasiswa"
	"be13/ca/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MahasiswaDelivery struct {
	mahasiswaService mahasiswa.ServiceInterface
}

func New(service mahasiswa.ServiceInterface, e *echo.Echo) {
	handler := &MahasiswaDelivery{
		mahasiswaService: service,
	}

	e.POST("/mahasiswa", handler.Create)
	e.PUT("/mahasiswa/:id", handler.Update)
	e.DELETE("/mahasiswa/:id", handler.Delete)
	e.GET("/mahasiswa/:id", handler.Read)
}

func (delivery *MahasiswaDelivery) Create(c echo.Context) error {
	userInput := MahasiswaRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	dataCore := requestToCore(userInput)
	err := delivery.mahasiswaService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create new users"))
}

func (delivery *MahasiswaDelivery) Update(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	userInput := MahasiswaRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := requestToCore(userInput)
	errUpt := delivery.mahasiswaService.Update(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (delivery *MahasiswaDelivery) Delete(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	errDel := delivery.mahasiswaService.Delete(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete user"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}

func (delivery *MahasiswaDelivery) Read(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	results, err := delivery.mahasiswaService.Read(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data", dataResponse))
}
