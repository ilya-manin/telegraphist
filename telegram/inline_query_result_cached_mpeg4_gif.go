package telegram

/*
InlineQueryResultCachedMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
*/
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`                            // Type of the result, must be mpeg4_gif
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Mpeg4FileID         string                `json:"mpeg4_file_id"`                   // A valid file identifier for the MP4 file
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video animation
}
