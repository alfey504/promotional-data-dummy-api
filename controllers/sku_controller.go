package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetSKUController(c *gin.Context) {
	skus := repository.GetSkus()
	c.JSON(http.StatusOK, models.APIRespone[[]models.SKU]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    skus,
	})
}
