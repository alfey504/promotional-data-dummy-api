package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/repository"
	"github.com/gin-gonic/gin"
)

func GetFulfillmentController(c *gin.Context) {
	maxPages := repository.GetMaxFulfillmentPages()

	pageQuery := c.Query("page")
	if pageQuery == "" {
		c.JSON(http.StatusBadRequest, models.PagedAPIResponse[any]{
			Status:  http.StatusBadRequest,
			Message: "the request must include the query'page' in the url",
			Page:    nil,
			MaxPage: maxPages,
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
			MaxPage: maxPages,
			Data:    nil,
		})
		c.Abort()
		return
	}

	if page < 1 || page > maxPages {
		c.JSON(http.StatusBadRequest, models.PagedAPIResponse[any]{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("the page number must be between %d and %d", 1, maxPages),
			Page:    nil,
			MaxPage: maxPages,
			Data:    nil,
		})
		c.Abort()
		return
	}

	fulfillmentHistory := repository.GetFulfillmentHistory(page)
	c.JSON(http.StatusOK, models.PagedAPIResponse[[]models.FulfillmentHistory]{
		Status:  http.StatusOK,
		Message: "Success",
		Page:    &page,
		MaxPage: maxPages,
		Data:    fulfillmentHistory,
	})
}
