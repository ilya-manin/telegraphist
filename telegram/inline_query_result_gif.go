package telegram

/*
InlineQueryResultGif represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
*/
type InlineQueryResultGif struct {
	Type                string                `json:"type"`                            // Type of the result, must be gif
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	GifURL              string                `json:"gif_url"`                         // A valid URL for the GIF file. File size must not exceed 1MB
	GifWidth            int                   `json:"gif_width,omitempty"`             // Optional. Width of the GIF
	GifHeight           int                   `json:"gif_height,omitempty"`            // Optional. Height of the GIF
	GifDuration         int                   `json:"gif_duration,omitempty"`          // Optional. Duration of the GIF
	ThumbURL            string                `json:"thumb_url"`                       // URL of the static thumbnail for the result (jpeg or gif)
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the GIF file to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the GIF animation
}
