package repository

import (
	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

const PRODUCT_REPO_PATH = "data\\products.json"
const SKU_REPO_PATH = "data\\skus.json"
const CUSTOMER_REPO_PATH = "data\\customers.json"
const BUNDLES_REPO_PATH = "data\\bundles.json"
const PROMOTION_REPO_PATH = "data\\promotions.json"
const SALES_REPO_PATH = "data\\sales.json"

var products = []models.Product{}
var skus = []models.SKU{}
var customers = []models.Customer{}
var bundles = []models.Bundle{}
var promotions = []models.Promotion{}
var sales = []models.Sale{}

func LoadData() error {
	if err := utils.DecodeJSON(PRODUCT_REPO_PATH, &products); err != nil {
		return err
	}

	if err := utils.DecodeJSON(SKU_REPO_PATH, &skus); err != nil {
		return err
	}

	if err := utils.DecodeJSON(CUSTOMER_REPO_PATH, &customers); err != nil {
		return err
	}

	if err := utils.DecodeJSON(BUNDLES_REPO_PATH, &bundles); err != nil {
		return err
	}

	if err := utils.DecodeJSON(PROMOTION_REPO_PATH, &promotions); err != nil {
		return err
	}

	if err := utils.DecodeJSON(SALES_REPO_PATH, &sales); err != nil {
		return err
	}

	return nil
}

func GetProducts() []models.Product {
	return products
}

func GetSkus() []models.SKU {
	return skus
}

func GetCustomers() []models.Customer {
	return customers
}

func GetBundles() []models.Bundle {
	return bundles
}

func GetPromotion() []models.Promotion {
	return promotions
}

func GetSales(page int) []models.Sale {
	pageDataLimit := 1000
	maxCount := len(sales) //25282

	pageStart := pageDataLimit * (page - 1)
	pageEnd := pageStart + pageDataLimit
	if pageEnd > maxCount {
		pageEnd = maxCount
	}
	return sales[pageStart:pageEnd]
}

func GetMaxSalesPages() int {
	maxCount := len(sales)
	pageDataLimit := 1000
	pageCount := maxCount / pageDataLimit
	if pageCount%pageDataLimit > 0 {
		pageCount++
	}
	return pageCount
}
