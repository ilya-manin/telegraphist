package telegram

type StopPollParams struct {
	ChatID      interface{}           `json:"chat_id"`                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID   int                   `json:"message_id"`             // IDentifier of the original message with the poll
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // A JSON-serialized object for a new message inline keyboard.
}

/*
Use this method to stop a poll which was sent by the bot. On success, the stopped Poll with the final results is returned.
*/
func (c *Client) StopPoll(params *StopPollParams) (Poll, error) {
	var err error
	var result Poll

	err = c.performRequest("stopPoll", params, nil, &result)
	return result, err
}
