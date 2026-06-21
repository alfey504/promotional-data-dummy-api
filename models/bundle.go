package models

type Bundle struct {
	BundleID          int     `json:"bundle_id"`
	BundleName        string  `json:"bundle_name"`
	BundleDescription string  `json:"bundle_description"`
	BundlePrice       float64 `json:"bundle_price"`
	SKUIDs            []int   `json:"sku_ids"`
}
