package models

import "dummyretaildata.com/dummydataserver/custom_types"

type FullfillmentHistory struct {
	FullfillmentID   int               `json:"fullfillment_id"`
	SkuID            int               `json:"sku_id"`
	FullfillmentDate custom_types.Date `json:"fullfillment_date"`
	QuantityReceived int               `json:"quantity_received"`
}
