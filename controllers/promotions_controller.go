package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetPromotionControllers(c *gin.Context) {
	promotions := repository.GetPromotion()
	c.JSON(http.StatusOK, models.APIRespone[[]models.Promotion]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    promotions,
	})
}
