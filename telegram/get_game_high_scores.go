package telegram

type GetGameHighScoresParams struct {
	UserID          int    `json:"user_id"`                     // Target user id
	ChatID          int64  `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID       int    `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. IDentifier of the sent message
	InlineMessageID string `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. IDentifier of the inline message
}

/*
Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
*/
func (c *Client) GetGameHighScores(params *GetGameHighScoresParams) ([]GameHighScore, error) {
	var err error
	var result []GameHighScore

	err = c.performRequest("getGameHighScores", params, nil, &result)

	return result, err
}
