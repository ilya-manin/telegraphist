package telegram

type SetGameScoreParams struct {
	UserID             int    `json:"user_id"`                        // User identifier
	Score              int    `json:"score"`                          // New score, must be non-negative
	Force              bool   `json:"force,omitempty"`                // Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"` // Pass True, if the game message should not be automatically edited to include the current scoreboard
	ChatID             int64  `json:"chat_id,omitempty"`              // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID          int    `json:"message_id,omitempty"`           // Required if inline_message_id is not specified. IDentifier of the sent message
	InlineMessageID    string `json:"inline_message_id,omitempty"`    // Required if chat_id and message_id are not specified. IDentifier of the inline message
}

/*
Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
*/
func (c *Client) SetGameScore(params *SetGameScoreParams) (interface{}, error) {
	var err error
	var result interface{}

	err = c.performRequest("setGameScore", params, nil, &result)

	switch v := result.(type) {
	case Message:
		result = v
	case bool:
		result = v
	}

	return result, err
}
