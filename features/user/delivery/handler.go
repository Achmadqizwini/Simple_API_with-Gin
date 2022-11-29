package delivery

import (
	"be13/ca/features/user"
	"be13/ca/middlewares"
	"be13/ca/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}

	e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/users", handler.Create)
	e.GET("/users/:id", handler.GetById)
	e.PUT("/users/:id", handler.UpdateData)
	e.DELETE("/users/:id", handler.DeleteUser)
}

func (delivery *UserDelivery) GetAll(c echo.Context) error {
	results, err := delivery.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}

func (delivery *UserDelivery) Create(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	dataCore := requestToCore(userInput)
	err := delivery.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create new users"))
}

func (delivery *UserDelivery) GetById(c echo.Context) error {
	id, errconv := strconv.Atoi(c.Param("id"))
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errconv.Error()))
	}

	IdUser, err := delivery.userService.GetById(id)
	dataResponse := fromCore(IdUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Get users "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
}

func (delivery *UserDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := requestToCore(userInput)
	errUpt := delivery.userService.UpdateUser(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (delivery *UserDelivery) DeleteUser(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	errDel := delivery.userService.DeleteUser(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete user"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
