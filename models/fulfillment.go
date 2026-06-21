package models

import "time"

type FulfillmentHistory struct {
	FulfillmentID    int       `json:"fulfillment_id"`
	SKUID            int       `json:"sku_id"`
	FulfillmentDate  time.Time `json:"fulfillment_date"`
	QuantityReceived int       `json:"quantity_received"`
}
