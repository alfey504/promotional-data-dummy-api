package models

type Customer struct {
	CustomerID     int    `json:"customer_id"`
	CustomerAge    int    `json:"customer_age"`
	CustomerGender string `json:"customer_gender"`
	Ethnicity      string `json:"ethnicity"`
}
