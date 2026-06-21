package controllers

import (
	"net/http"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetPromotionControllers(c *gin.Context) {
	promotions, err := repository.GetPromotion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIRespone[any]{
			Status:  http.StatusInternalServerError,
			Message: "There was an internal server issue while fetching your data",
			Data:    nil,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.APIRespone[[]models.Promotion]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    promotions,
	})
}
