package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetCustomersController(c *gin.Context) {
	customers := repository.GetCustomers()
	c.JSON(http.StatusOK, models.APIRespone[[]models.Customer]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    customers,
	})
}
