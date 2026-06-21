package repository

import "dummyretaildata.com/dummydataserver/models"

var maxSalesPages int

const salesPageDataLimit = 1000

func GetSales(page int) []models.Sale {
	maxCount := len(sales) //25282

	pageStart := salesPageDataLimit * (page - 1)
	pageEnd := pageStart + salesPageDataLimit
	if pageEnd > maxCount {
		pageEnd = maxCount
	}
	return sales[pageStart:pageEnd]
}

func findMaxSalePages() int {
	maxCount := len(sales)
	pageCount := maxCount / salesPageDataLimit
	if pageCount%salesPageDataLimit > 0 {
		pageCount++
	}
	return pageCount
}

func GetMaxSalesPages() int {
	return maxSalesPages
}
