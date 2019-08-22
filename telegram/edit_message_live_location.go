package telegram

type EditMessageLiveLocationParams struct {
	ChatID          interface{}           `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       int                   `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. IDentifier of the message to edit
	InlineMessageID string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. IDentifier of the inline message
	Latitude        float64               `json:"latitude"`                    // Latitude of new location
	Longitude       float64               `json:"longitude"`                   // Longitude of new location
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for a new inline keyboard.
}

/*
Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (c *Client) EditMessageLiveLocation(params *EditMessageLiveLocationParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("editMessageLiveLocation", params, nil, &result)

	return result, err
}
