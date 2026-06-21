package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetBundlesController(c *gin.Context) {
	bundles := repository.GetBundles()
	c.JSON(http.StatusOK, models.APIRespone[[]models.Bundle]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    bundles,
	})
}
