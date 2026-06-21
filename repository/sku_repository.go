package repository

import (
	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

const SKU_REPO_PATH = "data\\skus.json"

func GetSkus() ([]models.Sku, error) {
	skus := []models.Sku{}
	if err := utils.DecodeJSON(SKU_REPO_PATH, &skus); err != nil {
		return nil, err
	}
	return skus, nil
}
