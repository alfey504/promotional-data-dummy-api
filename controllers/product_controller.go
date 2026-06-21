package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetProductController(c *gin.Context) {
	products := repository.GetProducts()
	c.JSON(http.StatusOK, models.APIRespone[[]models.Product]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    products,
	})
}
