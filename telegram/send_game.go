package telegram

type SendGameParams struct {
	ChatID              int64                 `json:"chat_id"`                        // Unique identifier for the target chat
	GameShortName       string                `json:"game_short_name"`                // Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
	DisableNotification bool                  `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int                   `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`         // A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.
}

/*
Use this method to send a game. On success, the sent Message is returned.
*/
func (c *Client) SendGame(params *SendGameParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("sendGame", params, nil, &result)

	return result, err
}
