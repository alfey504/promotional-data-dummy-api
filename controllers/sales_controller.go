package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetSalesController(c *gin.Context) {

	maxSalesPages := repository.GetMaxSalesPages()

	pageQuery := c.Query("page")
	if pageQuery == "" {
		c.JSON(http.StatusBadRequest, models.PagedAPIResponse[any]{
			Status:  http.StatusBadRequest,
			Message: "the request must include the query'page' in the url",
			Page:    nil,
			MaxPage: maxSalesPages,
			Data:    nil,
		})
		c.Abort()
		return
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.PagedAPIResponse[any]{
			Status:  http.StatusBadRequest,
			Message: "the query page must be integer number",
			Page:    nil,
			MaxPage: maxSalesPages,
			Data:    nil,
		})
		c.Abort()
		return
	}

	if page < 1 || page > maxSalesPages {
		c.JSON(http.StatusBadRequest, models.PagedAPIResponse[any]{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("the page number must be between %d and %d", 1, maxSalesPages),
			Page:    nil,
			MaxPage: maxSalesPages,
			Data:    nil,
		})
		c.Abort()
		return
	}

	sales := repository.GetSales(page)
	c.JSON(http.StatusOK, models.PagedAPIResponse[[]models.Sale]{
		Status:  http.StatusOK,
		Message: "Success",
		Page:    &page,
		MaxPage: maxSalesPages,
		Data:    sales,
	})
}
