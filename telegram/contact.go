package telegram

/*
Contact represents a phone contact.
*/
type Contact struct {
	PhoneNumber string `json:"phone_number"`        // Contact's phone number
	FirstName   string `json:"first_name"`          // Contact's first name
	LastName    string `json:"last_name,omitempty"` // Optional. Contact's last name
	UserID      int    `json:"user_id,omitempty"`   // Optional. Contact's user identifier in Telegram
	Vcard       string `json:"vcard,omitempty"`     // Optional. Additional data about the contact in the form of a vCard
}
