package models

import (
	"dummyretaildata.com/dummydataserver/custom_types"
)

type Sale struct {
	SalesID      int               `json:"sales_id"`
	CustomerID   int               `json:"customer_id"`
	PromotionIDs []int             `json:"promotion_ids"`
	SaleDate     custom_types.Date `json:"sale_date"`
	FinalPrice   float64           `json:"final_price"`
	SKUSales     []SKUSale         `json:"sku_sales"`
	BundleSales  []BundleSale      `json:"bundle_sales"`
}

type BundleSale struct {
	BundleID int `json:"bundle_id"`
	Quantity int `json:"quantity"`
}

type SKUSale struct {
	SKUID    int `json:"sku_id"`
	Quantity int `json:"quantity"`
}
