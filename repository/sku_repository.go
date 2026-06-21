package repository

import (
	"fmt"

	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

const SKU_REPO_PATH = "data\\skus.json"

func GetSkus() ([]models.SKU, error) {
	skus := []models.SKU{}
	if err := utils.DecodeJSON(SKU_REPO_PATH, &skus); err != nil {
		fmt.Printf("Error : %s", err)
		return nil, err
	}
	return skus, nil
}
