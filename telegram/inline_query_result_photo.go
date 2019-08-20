package telegram

/*
InlineQueryResultPhoto represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
*/
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`                            // Type of the result, must be photo
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	PhotoURL            string                `json:"photo_url"`                       // A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	ThumbURL            string                `json:"thumb_url"`                       // URL of the thumbnail for the photo
	PhotoWidth          int64                 `json:"photo_width,omitempty"`           // Optional. Width of the photo
	PhotoHeight         int64                 `json:"photo_height,omitempty"`          // Optional. Height of the photo
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the photo to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the photo
}
