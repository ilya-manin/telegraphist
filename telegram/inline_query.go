package telegram

/*
InlineQuery represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
*/
type InlineQuery struct {
	ID       string    `json:"id"`                 // Unique identifier for this query
	From     *User     `json:"from"`               // Sender
	Location *Location `json:"location,omitempty"` // Optional. Sender location, only for bots that request user location
	Query    string    `json:"query"`              // Text of the query (up to 512 characters)
	Offset   string    `json:"offset"`             // Offset of the results to be returned, can be controlled by the bot
}
