package telegram

/*
ShippingOption represents one shipping option.
*/
type ShippingOption struct {
	ID     string          `json:"id"`     // Shipping option identifier
	Title  string          `json:"title"`  // Option title
	Prices []*LabeledPrice `json:"prices"` // List of price portions
}
