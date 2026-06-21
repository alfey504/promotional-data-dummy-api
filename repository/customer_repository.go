package repository

import (
	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

const CUSTOMER_REPO_PATH = "data\\customers.json"

func GetCustomers() ([]models.Customer, error) {
	customers := []models.Customer{}
	if err := utils.DecodeJSON(CUSTOMER_REPO_PATH, &customers); err != nil {
		return nil, err
	}
	return customers, nil
}
