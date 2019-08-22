package telegram

/*
InlineQueryResultVenue represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
*/
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`                            // Type of the result, must be venue
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	Latitude            float64               `json:"latitude"`                        // Latitude of the venue location in degrees
	Longitude           float64               `json:"longitude"`                       // Longitude of the venue location in degrees
	Title               string                `json:"title"`                           // Title of the venue
	Address             string                `json:"address"`                         // Address of the venue
	FoursquareID        string                `json:"foursquare_id,omitempty"`         // Optional. Foursquare identifier of the venue if known
	FoursquareType      string                `json:"foursquare_type,omitempty"`       // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the venue
	ThumbURL            string                `json:"thumb_url,omitempty"`             // Optional. URL of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}
