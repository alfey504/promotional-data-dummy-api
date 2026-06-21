package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetSKUController(c *gin.Context) {
	skus, err := repository.GetSkus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIRespone[any]{
			Status:  http.StatusInternalServerError,
			Message: "There was an internal server error while fetching your data",
			Data:    nil,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, models.APIRespone[[]models.Sku]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    skus,
	})
}
