package telegram

/*
CallbackGame is a placeholder, currently holds no information. Use BotFather to set up your game.
*/
type CallbackGame struct {
	UserID             int64  `json:"user_id"`                     // User identifier
	Score              int64  `json:"score"`                       // New score, must be non-negative
	Force              bool   `json:"force"`                       // Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message"`        // Pass True, if the game message should not be automatically edited to include the current scoreboard
	ChatID             int64  `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID          int64  `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. IDentifier of the sent message
	InlineMessageID    string `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. IDentifier of the inline message
}
