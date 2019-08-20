package telegram

/*
InlineQueryResultDocument represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
*/
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`                            // Type of the result, must be document
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                           // Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DocumentURL         string                `json:"document_url"`                    // A valid URL for the file
	MimeType            string                `json:"mime_type"`                       // Mime type of the content of the file, either “application/pdf” or “application/zip”
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
	ThumbURL            string                `json:"thumb_url,omitempty"`             // Optional. URL of the thumbnail (jpeg only) for the file
	ThumbWidth          int64                 `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int64                 `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}
