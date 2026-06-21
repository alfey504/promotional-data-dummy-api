package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetSalesController(c *gin.Context) {
	sales := repository.GetSales()
	c.JSON(http.StatusOK, models.APIRespone[[]models.Sale]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    sales,
	})
}
