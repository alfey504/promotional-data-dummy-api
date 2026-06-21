package models

import (
	"dummyretaildata.com/dummydataserver/custom_types"
)

type SKU struct {
	SKUID               int                `json:"sku_id"`
	ProductID           int                `json:"product_id"`
	SKUName             string             `json:"sku_name"`
	Size                string             `json:"size"`
	Color               string             `json:"color"`
	LastFulfillmentDate *custom_types.Date `json:"last_fulfillment_date"`
	NextFulfillmentDate *custom_types.Date `json:"next_fulfillment_date"`
	InStock             int                `json:"in_stock"`
	Price               float64            `json:"price"`
}
