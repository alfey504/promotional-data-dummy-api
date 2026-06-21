package models

import (
	"dummyretaildata.com/dummydataserver/custom_types"
)

type Promotion struct {
	PromotionID          int               `json:"promotion_id"`
	PromotionName        string            `json:"promotion_name"`
	PromotionDescription string            `json:"promotion_description"`
	PromotionType        string            `json:"promotion_type"`
	TargetBundleIDs      []int             `json:"target_bundle_ids"`
	TargetSKUIDs         []int             `json:"target_sku_ids"`
	DiscountPercent      int               `json:"discount_percent"`
	StartDate            custom_types.Date `json:"start_date"`
	EndDate              custom_types.Date `json:"end_date"`
}
