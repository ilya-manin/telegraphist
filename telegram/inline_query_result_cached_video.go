package telegram

/*
InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
*/
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`                            // Type of the result, must be video
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VideoFileID         string                `json:"video_file_id"`                   // A valid file identifier for the video file
	Title               string                `json:"title"`                           // Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the video to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video
}
