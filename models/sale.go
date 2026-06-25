package models

import "dummyretaildata.com/dummydataserver/custom_types"

type Sale struct {
	SalesID      int               `json:"sales_id"`
	SkuID        int               `json:"sku_id"`
	PromotionID  *int              `json:"promotion_id"`
	RegularPrice float64           `json:"regular_price"`
	FinalPrice   float64           `json:"final_price"`
	CustomerID   *int              `json:"customer_id"`
	SaleDate     custom_types.Date `json:"sale_date"`
}
