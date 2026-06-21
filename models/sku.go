package models

type Sku struct {
	SkuID     int    `json:"sku_id"`
	ProductID int    `json:"product_id"`
	SkuName   string `json:"sku_name"`
	Size      string `json:"size"`
	Color     string `json:"color"`
}
