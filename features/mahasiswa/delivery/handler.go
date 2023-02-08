package delivery

import (
	"be13/ca/features/mahasiswa"
	"be13/ca/utils/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MahasiswaDelivery struct {
	mahasiswaService mahasiswa.ServiceInterface
}

func New(service mahasiswa.ServiceInterface, e *gin.Engine) {
	handler := &MahasiswaDelivery{
		mahasiswaService: service,
	}
	e.POST("/mahasiswa", handler.Create)
	e.PUT("/mahasiswa/:id", handler.Update)
	e.DELETE("/mahasiswa/:id", handler.Delete)
	e.GET("/mahasiswa/:id", handler.Read)
	e.GET("/mahasiswa", handler.GetAll)
}

func (delivery *MahasiswaDelivery) Create(c *gin.Context) {
	userInput := mahasiswa.Core{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	err := delivery.mahasiswaService.Create(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	c.IndentedJSON(http.StatusOK, helper.SuccessResponse("success create new users"))
}

func (delivery *MahasiswaDelivery) Update(c *gin.Context) {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	userInput := mahasiswa.Core{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	errUpt := delivery.mahasiswaService.Update(userInput, id)
	if errUpt != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (delivery *MahasiswaDelivery) Delete(c *gin.Context) {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	errDel := delivery.mahasiswaService.Delete(id)
	if errDel != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete user"+errDel.Error()))
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}

func (delivery *MahasiswaDelivery) Read(c *gin.Context) {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	results, err := delivery.mahasiswaService.Read(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data", results))
}

func (delivery *MahasiswaDelivery) GetAll(c *gin.Context) {
	result, err := delivery.mahasiswaService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "error get all mahasiswa",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success get all data",
		"data":    result,
	})
}
