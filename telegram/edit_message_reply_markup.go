package telegram

type EditMessageReplyMarkupParams struct {
	ChatID          interface{}           `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       int                   `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. IDentifier of the message to edit
	InlineMessageID string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. IDentifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for an inline keyboard.
}

/*
Use this method to edit only the reply markup of messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (c *Client) EditMessageReplyMarkup(params *EditMessageReplyMarkupParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("editMessageReplyMarkup", params, nil, &result)

	return result, err
}
