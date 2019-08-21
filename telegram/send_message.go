package telegram

type SendMessageParams struct {
	ChatID                interface{} `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                  string      `json:"text"`                               // Text of the message to be sent
	ParseMode             string      `json:"parse_mode,omitempty"`               // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"` // Disables link previews for links in this message
	DisableNotification   bool        `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID      int64       `json:"reply_to_message_id,omitempty"`      // If the message is a reply, ID of the original message
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send text messages. On success, the sent Message is returned.
*/
func (c *Client) SendMessage(params *SendMessageParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("sendMessage", params, nil, &result)

	return result, err
}
