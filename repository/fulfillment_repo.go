package repository

import "dummyretaildata.com/dummydataserver/models"

var maxFulfillmentPage int

const fulfillmentPageDataLimit = 1000

func GetFulfillmentHistory(page int) []models.FulfillmentHistory {
	maxCount := len(sales)
	pageStart := fulfillmentPageDataLimit * (page - 1)
	pageEnd := pageStart + fulfillmentPageDataLimit
	if pageEnd > maxCount {
		pageEnd = maxCount
	}
	return fulfillment_history[pageStart:pageEnd]
}

func findMaxFulfillmentPages() int {
	maxCount := len(fulfillment_history)
	pageCount := maxCount / fulfillmentPageDataLimit
	if pageCount%fulfillmentPageDataLimit > 0 {
		pageCount++
	}
	return pageCount
}

func GetMaxFulfillmentPages() int {
	return maxFulfillmentPage
}
