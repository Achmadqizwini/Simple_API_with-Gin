package delivery

import (
	"be13/ca/features/book"
	"be13/ca/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookDelivery struct {
	bookService book.ServiceInterface
}

func New(service book.ServiceInterface, e *echo.Echo) {
	handler := &BookDelivery{
		bookService: service,
	}

	e.GET("/books", handler.GetAll)
	e.POST("/books", handler.Create)
	e.GET("/books/:id", handler.GetById)
	e.PUT("/books/:id", handler.UpdateData)
	e.DELETE("/books/:id", handler.DeleteBook)
}

func (delivery *BookDelivery) GetAll(c echo.Context) error {
	results, err := delivery.bookService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data books"))
	}

	dataResponse := ListModelBook(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all books", dataResponse))
}

func (delivery *BookDelivery) Create(c echo.Context) error {
	bookInput := BookRequest{}
	errBind := c.Bind(&bookInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	dataCore := BookRequestToCore(bookInput)
	err := delivery.bookService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create new books"))
}

func (delivery *BookDelivery) GetById(c echo.Context) error {
	id, errconv := strconv.Atoi(c.Param("id"))
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errconv.Error()))
	}

	IdBook, err := delivery.bookService.GetById(id)
	dataResponse := DataIDBookRespon(IdBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Get books "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get books", dataResponse))
}

func (delivery BookDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	bookInput := BookRequest{}
	errBind := c.Bind(&bookInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := BookRequestToCore(bookInput)
	errUpt := delivery.bookService.UpdateBook(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (delivery *BookDelivery) DeleteBook(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	errDel := delivery.bookService.DeleteBook(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete book"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
