package models

import "time"

type SKU struct {
	SKUID               int        `json:"sku_id"`
	ProductID           int        `json:"product_id"`
	SKUName             string     `json:"sku_name"`
	Size                string     `json:"size"`
	Color               string     `json:"color"`
	LastFulfillmentDate *time.Time `json:"last_fulfillment_date,omitempty"`
	NextFulfillmentDate *time.Time `json:"next_fulfillment_date,omitempty"`
	InStock             int        `json:"in_stock"`
	Price               float64    `json:"price"`
}
