package repository

import (
	"fmt"
	"time"

	"dummyretaildata.com/dummydataserver/custom_types"
	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

func GetPromotion() ([]models.Promotion, error) {

	promotionNames := []string{
		"Summer Splash Savings",
		"Weekend Flash Sale",
		"Mid-Season Clearance",
		"First Order Reward",
		"Holiday Mega Deal",
		"VIP Exclusive Discount",
		"Back to School Special",
		"Black Friday Doorbuster",
		"Anniversary Celebration",
		"Payday Treat",
	}

	promotionDescriptions := []string{
		"Beat the heat with refreshing discounts on all our top-tier summer products.",
		"Don't blink! Get massive price drops on selected items this weekend only.",
		"Making room for new arrivals. Snag your favorites before they are gone forever.",
		"A special welcome gift just for you—enjoy a discount on your very first purchase.",
		"Spread the festive cheer with our biggest savings event of the holiday season.",
		"An exclusive thank-you offer reserved only for our most loyal reward members.",
		"Stock up on all your daily essentials and get completely ready for the new semester.",
		"The ultimate shopping event with unprecedented price cuts for a limited time.",
		"We are celebrating another year of excellence by passing the savings on to you.",
		"You worked hard this month! Treat yourself to something special with these deals.",
	}

	promotionTypes := []string{
		"discount_sale",
		"bundle_sale",
		"bogo",
	}
	promotions := []models.Promotion{}

	count := 350
	for i := range count {

		promotionIndex := utils.RandomIntBetween(0, len(promotionNames)-1)
		promotionType := promotionTypes[utils.RandomIntBetween(0, len(promotionTypes)-1)]

		targetBundles := []int{}
		targetSkUs := []int{}

		if promotionType == "bundle_sale" {
			targetBundles = utils.SampleRandInt(1, 99, utils.RandomIntBetween(4, 10))
		} else {
			targetSkUs = utils.SampleRandInt(1, 167, utils.RandomIntBetween(10, 20))
		}

		layout := "02-01-2006"
		promotionRangeStart, err := time.Parse(layout, "01-01-2023")
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil, err
		}
		promotionRangeEnd := time.Now()

		startDate := utils.RandomDate(promotionRangeStart, promotionRangeEnd)
		endDate := startDate.AddDate(0, 0, utils.RandomIntBetween(14, 30))

		promtion := models.Promotion{
			PromotionID:          i,
			PromotionName:        promotionNames[promotionIndex],
			PromotionType:        promotionType,
			PromotionDescription: promotionDescriptions[promotionIndex],
			TargetBundles:        targetBundles,
			TargetSKUs:           targetSkUs,
			DiscountPercent:      0,
			StartDate:            custom_types.Date{Time: startDate},
			EndDate:              custom_types.Date{Time: endDate},
		}
		promotions = append(promotions, promtion)
	}

	return promotions, nil
}
