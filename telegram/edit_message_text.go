package telegram

type EditMessageTextParams struct {
	ChatID                interface{}           `json:"chat_id,omitempty"`                  // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID             int                   `json:"message_id,omitempty"`               // Required if inline_message_id is not specified. IDentifier of the message to edit
	InlineMessageID       string                `json:"inline_message_id,omitempty"`        // Required if chat_id and message_id are not specified. IDentifier of the inline message
	Text                  string                `json:"text"`                               // New text of the message
	ParseMode             string                `json:"parse_mode,omitempty"`               // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"` // Disables link previews for links in this message
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // A JSON-serialized object for an inline keyboard.
}

/*
Use this method to edit text and game messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (c *Client) EditMessageText(params *EditMessageTextParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("editMessageText", params, nil, &result)

	return result, err
}
