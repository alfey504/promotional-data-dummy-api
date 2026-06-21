package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetSalesController(c *gin.Context) {
	sales, err := repository.GetSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIRespone[any]{
			Status:  http.StatusInternalServerError,
			Message: "There was an internal server error while fetching your data",
			Data:    nil,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, models.APIRespone[[]models.Sale]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    sales,
	})
}
