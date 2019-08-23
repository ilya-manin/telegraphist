package telegram

type SendInvoiceParams struct {
	ChatID                    int64                 `json:"chat_id"`                                 // Unique identifier for the target private chat
	Title                     string                `json:"title"`                                   // Product name, 1-32 characters
	Description               string                `json:"description"`                             // Product description, 1-255 characters
	Payload                   string                `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string                `json:"provider_token"`                          // Payments provider token, obtained via Botfather
	StartParameter            string                `json:"start_parameter"`                         // Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
	Currency                  string                `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    *[]LabeledPrice       `json:"prices"`                                  // Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	ProviderData              string                `json:"provider_data,omitempty"`                 // JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoURL                  string                `json:"photo_url,omitempty"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int                   `json:"photo_size,omitempty"`                    // Photo size
	PhotoWidth                int                   `json:"photo_width,omitempty"`                   // Photo width
	PhotoHeight               int                   `json:"photo_height,omitempty"`                  // Photo height
	NeedName                  bool                  `json:"need_name,omitempty"`                     // Pass True, if you require the user's full name to complete the order
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`             // Pass True, if you require the user's phone number to complete the order
	NeedEmail                 bool                  `json:"need_email,omitempty"`                    // Pass True, if you require the user's email address to complete the order
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`         // Pass True, if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"` // Pass True, if user's phone number should be sent to provider
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`        // Pass True, if user's email address should be sent to provider
	IsFlexible                bool                  `json:"is_flexible,omitempty"`                   // Pass True, if the final price depends on the shipping method
	DisableNotification       bool                  `json:"disable_notification,omitempty"`          // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID          int                   `json:"reply_to_message_id,omitempty"`           // If the message is a reply, ID of the original message
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

/*
Use this method to send invoices. On success, the sent Message is returned.
*/
func (c *Client) SendInvoice(params *SendInvoiceParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("sendInvoice", params, nil, &result)

	return result, err
}
