package models

import (
	"dummyretaildata.com/dummydataserver/custom_types"
)

type PromotionType string

const (
	PercentageOff PromotionType = "Percentage Off"
	BOGO          PromotionType = "BOGO"
	FlashSale     PromotionType = "Flash Sale"
	SeasonalSale  PromotionType = "Seasonal Sale"
)

type Promotion struct {
	PromotionID          int               `json:"promotion_id"`
	PromotionName        string            `json:"promotion_name"`
	PromotionDescription *string           `json:"promotion_description"`
	PromotionType        PromotionType     `json:"promotion_type"`
	DiscountPercent      *int              `json:"discount_percent"`
	StartDate            custom_types.Date `json:"start_date"`
	EndDate              custom_types.Date `json:"end_date"`
	TargetSkus           []int             `json:"target_skus"`
}
