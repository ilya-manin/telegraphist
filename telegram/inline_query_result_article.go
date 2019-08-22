package telegram

/*
InlineQueryResultArticle represents a link to an article or web page.
*/
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                   // Type of the result, must be article
	ID                  string                `json:"id"`                     // Unique identifier for this result, 1-64 Bytes
	Title               string                `json:"title"`                  // Title of the result
	InputMessageContent *InputMessageContent  `json:"input_message_content"`  // Content of the message to be sent
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. Inline keyboard attached to the message
	URL                 string                `json:"url,omitempty"`          // Optional. URL of the result
	HideURL             bool                  `json:"hide_url,omitempty"`     // Optional. Pass True, if you don't want the URL to be shown in the message
	Description         string                `json:"description,omitempty"`  // Optional. Short description of the result
	ThumbURL            string                `json:"thumb_url,omitempty"`    // Optional. URL of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width,omitempty"`  // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"` // Optional. Thumbnail height
}
