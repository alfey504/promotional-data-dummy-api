package models

type Product struct {
	ProductID          int     `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription *string `json:"product_description"`
	ProductBrand       string  `json:"product_brand"`
	ProductCategory    string  `json:"product_category"`
}
