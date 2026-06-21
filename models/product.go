package models

import "time"

type Product struct {
	ProductID           int       `json:"product_id"`
	ProductName         string    `json:"product_name"`
	ProductDescription  string    `json:"product_description"`
	ProductBrand        string    `json:"product_brand"`
	ProductCategory     string    `json:"product_category"`
	Stock               int       `json:"stock"`
	LastFulfillmentDate time.Time `json:"last_fulfillment_date"`
	NextFulfillmentDate time.Time `json:"next_fulfillment_date"`
}
