package telegram

type SendContactParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	PhoneNumber         string      `json:"phone_number"`                   // Contact's phone number
	FirstName           string      `json:"first_name"`                     // Contact's first name
	LastName            string      `json:"last_name,omitempty"`            // Contact's last name
	Vcard               string      `json:"vcard,omitempty"`                // Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
}

/*
Use this method to send phone contacts. On success, the sent Message is returned.
*/
func (c *Client) SendContact(params *SendContactParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("sendContact", params, nil, &result)

	return result, err
}
