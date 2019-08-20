package telegram

/*
OrderInfo represents information about an order.
*/
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`             // Optional. User name
	PhoneNumber     string           `json:"phone_number,omitempty"`     // Optional. User's phone number
	Email           string           `json:"email,omitempty"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` // Optional. User shipping address
}
