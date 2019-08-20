package telegram

/*
ShippingAddress represents a shipping address.
*/
type ShippingAddress struct {
	CountryCode string `json:"country_code"` // ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}
