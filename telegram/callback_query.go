package telegram

/*
CallbackQuery represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
*/
type CallbackQuery struct {
	ID              string   `json:"id"`                          // Unique identifier for this query
	From            *User    `json:"from"`                        // Sender
	Message         *Message `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageID string   `json:"inline_message_id,omitempty"` // Optional. IDentifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string   `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string   `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	GameShortName   string   `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}
