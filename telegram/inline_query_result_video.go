package telegram

/*
InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
*/
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`                            // Type of the result, must be video
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VideoURL            string                `json:"video_url"`                       // A valid URL for the embedded video player or video file
	MimeType            string                `json:"mime_type"`                       // Mime type of the content of video url, “text/html” or “video/mp4”
	ThumbURL            string                `json:"thumb_url"`                       // URL of the thumbnail (jpeg only) for the video
	Title               string                `json:"title"`                           // Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the video to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	VideoWidth          int64                 `json:"video_width,omitempty"`           // Optional. Video width
	VideoHeight         int64                 `json:"video_height,omitempty"`          // Optional. Video height
	VideoDuration       int64                 `json:"video_duration,omitempty"`        // Optional. Video duration in seconds
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
}
