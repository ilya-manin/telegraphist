package telegram

/*
ShippingQuery contains information about an incoming shipping query.
*/
type ShippingQuery struct {
	ID              string           `json:"id"`               // Unique query identifier
	From            *User            `json:"from"`             // User who sent the query
	InvoicePayload  string           `json:"invoice_payload"`  // Bot specified invoice payload
	ShippingAddress *ShippingAddress `json:"shipping_address"` // User specified shipping address
}
