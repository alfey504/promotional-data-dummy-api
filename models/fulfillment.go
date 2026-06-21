package models

import (
	"dummyretaildata.com/dummydataserver/custom_types"
)

type FulfillmentHistory struct {
	FulfillmentID    int               `json:"fulfillment_id"`
	SKUID            int               `json:"sku_id"`
	FulfillmentDate  custom_types.Date `json:"fulfillment_date"`
	QuantityReceived int               `json:"quantity_received"`
}
