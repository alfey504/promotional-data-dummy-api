package repository

import (
	"dummyretaildata.com/dummydataserver/models"
	"dummyretaildata.com/dummydataserver/utils"
)

func GetBundles() ([]models.Bundle, error) {
	bundleNames := []string{
		"Everyday Essentials for Your Face",
		"Hydration Hit Facial Kit",
		"The Ultimate Glow Up Bundle",
		"Daily Defense & Shield",
		"Clear Skin Confidence Set",
		"Nighttime Renewal Regime",
		"Fresh Start Detox Collection",
		"Soothe & Calm Sensitive Pack",
		"Ageless Radiance Routine",
		"Bright & Awake Morning Trio",
	}

	bundleDescriptions := []string{
		"All the required stuff for your face to get through the day.",
		"A hard-hitting moisture boost for dry, dehydrated skin.",
		"Curated premium products to give your skin a natural, luminous glow.",
		"Essential SPF and antioxidant protection against daily environmental stressors.",
		"A powerful combination to target blemishes and keep your skin clear.",
		"Deeply nourishing formulas that repair and recharge your skin while you sleep.",
		"Clay masks and exfoliating scrubs designed to deeply purify your pores.",
		"Gentle, fragrance-free essentials to reduce redness and irritation.",
		"Advanced formulas designed to smooth fine lines and restore youthful bounce.",
		"Energy-boosting serums and eye creams to kickstart your morning routine.",
	}

	bundles := []models.Bundle{}

	count := 120
	for i := range count {
		target_sku_count := utils.RandomIntBetween(2, 6)
		bundle := models.Bundle{
			BundleID:          i,
			BundleName:        bundleNames[utils.RandomIntBetween(0, len(bundleNames))],
			BundleDescription: bundleDescriptions[utils.RandomIntBetween(0, len(bundleDescriptions))],
			BundlePrice:       float64(utils.RandomIntBetween(10, 30)) * float64(target_sku_count),
			SKUs:              utils.SampleRandInt(1, 165, target_sku_count),
		}
		bundles = append(bundles, bundle)
	}

	return bundles, nil
}
