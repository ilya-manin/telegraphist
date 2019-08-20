package telegram

/*
InlineQueryResultGame represents a Game.
*/
type InlineQueryResultGame struct {
	Type          string                `json:"type"`                   // Type of the result, must be game
	ID            string                `json:"id"`                     // Unique identifier for this result, 1-64 bytes
	GameShortName string                `json:"game_short_name"`        // Short name of the game
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. Inline keyboard attached to the message
}
