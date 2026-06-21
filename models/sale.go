package models

import (
	"dummyretaildata.com/dummydataserver/custom_types"
)

type Sale struct {
	SalesID     int               `json:"sales_id"`
	SKUs        []int             `json:"skus"`
	Bundles     []int             `json:"bundles"`
	PromotionID *int              `json:"promotion_id,omitempty"`
	FinalPrice  float64           `json:"final_price"`
	CustomerID  int               `json:"customer_id"`
	SaleDate    custom_types.Date `json:"sale_date"`
}
