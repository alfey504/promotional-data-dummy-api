package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetBundlesController(c *gin.Context) {
	bundles, err := repository.GetBundles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIRespone[any]{
			Status:  http.StatusInternalServerError,
			Message: "There was an internal server issue while fetching your data",
			Data:    nil,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, models.APIRespone[[]models.Bundle]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    bundles,
	})
}
