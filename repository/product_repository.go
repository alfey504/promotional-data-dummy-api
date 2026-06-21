package repository

import (
	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

const PRODUCT_REPO_PATH = "new_data\\products.json"

func GetProducts() ([]models.Product, error) {
	products := []models.Product{}
	if err := utils.DecodeJSON(PRODUCT_REPO_PATH, &products); err != nil {
		return nil, err
	}
	return products, nil
}
