package telegram

/*
InlineQueryResultLocation represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
*/
type InlineQueryResultLocation struct {
	Type                string                `json:"type"`                            // Type of the result, must be location
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	Latitude            float64               `json:"latitude"`                        // Location latitude in degrees
	Longitude           float64               `json:"longitude"`                       // Location longitude in degrees
	Title               string                `json:"title"`                           // Location title
	LivePeriod          int64                 `json:"live_period,omitempty"`           // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the location
	ThumbURL            string                `json:"thumb_url,omitempty"`             // Optional. URL of the thumbnail for the result
	ThumbWidth          int64                 `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int64                 `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}
