package telegram

type StopMessageLiveLocationParams struct {
	ChatID          interface{}           `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       int                   `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. IDentifier of the message with live location to stop
	InlineMessageID string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. IDentifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for a new inline keyboard.
}

/*
Use this method to stop updating a live location message before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
*/
func (c *Client) StopMessageLiveLocation(params *StopMessageLiveLocationParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("stopMessageLiveLocation", params, nil, &result)

	return result, err
}
