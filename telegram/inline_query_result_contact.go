package telegram

/*
InlineQueryResultContact represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
*/
type InlineQueryResultContact struct {
	Type                string                `json:"type"`                            // Type of the result, must be contact
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	PhoneNumber         string                `json:"phone_number"`                    // Contact's phone number
	FirstName           string                `json:"first_name"`                      // Contact's first name
	LastName            string                `json:"last_name,omitempty"`             // Optional. Contact's last name
	Vcard               string                `json:"vcard,omitempty"`                 // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the contact
	ThumbURL            string                `json:"thumb_url,omitempty"`             // Optional. URL of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}
