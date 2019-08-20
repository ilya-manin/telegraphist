package telegram

/*
InlineQueryResultAudio represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
*/
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`                            // Type of the result, must be audio
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	AudioURL            string                `json:"audio_url"`                       // A valid URL for the audio file
	Title               string                `json:"title"`                           // Title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Performer           string                `json:"performer,omitempty"`             // Optional. Performer
	AudioDuration       int64                 `json:"audio_duration,omitempty"`        // Optional. Audio duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}
