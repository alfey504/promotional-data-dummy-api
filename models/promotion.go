package models

import "time"

type Promotion struct {
	PromotionID     int       `json:"promotion_id"`
	PromotionName   string    `json:"promotion_name"`
	PromotionType   string    `json:"promotion_type"`
	TargetBundles   []int     `json:"target_bundles"`
	TargetSKUs      []int     `json:"target_skus"`
	DiscountPercent int       `json:"discount_percent"`
	StartDate       time.Time `json:"start_date"`
}
