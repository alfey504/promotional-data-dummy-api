package repository

import (
	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

const PRODUCT_REPO_PATH = "data/products.json"
const SKU_REPO_PATH = "data/skus.json"
const CUSTOMER_REPO_PATH = "data/customers.json"
const PROMOTION_REPO_PATH = "data/promotions.json"
const SALES_REPO_PATH = "data/sales.json"
const FULFILLMENT_REPO_PATH = "data/fulfillment_history.json"

var products = []models.Product{}
var skus = []models.Sku{}
var customers = []models.Customer{}
var promotions = []models.Promotion{}
var sales = []models.Sale{}
var fulfillment_history = []models.FullfillmentHistory{}

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

	if err := utils.DecodeJSON(PROMOTION_REPO_PATH, &promotions); err != nil {
		return err
	}

	if err := utils.DecodeJSON(SALES_REPO_PATH, &sales); err != nil {
		return err
	}

	maxSalesPages = findMaxSalePages()

	if err := utils.DecodeJSON(FULFILLMENT_REPO_PATH, &fulfillment_history); err != nil {
		return err
	}

	maxFulfillmentPage = findMaxFulfillmentPages()

	return nil
}

func GetProducts() []models.Product {
	return products
}

func GetSkus() []models.Sku {
	return skus
}

func GetCustomers() []models.Customer {
	return customers
}

func GetPromotion() []models.Promotion {
	return promotions
}
