package telegram

/*
InputMediaVideo represents a video to be sent.
*/
type InputMediaVideo struct {
	Type              string      `json:"type"`                         // Type of the result, must be video
	Media             interface{} `json:"media"`                        // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb             interface{} `json:"thumb,omitempty"`              // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption           string      `json:"caption,omitempty"`            // Optional. Caption of the video to be sent, 0-1024 characters
	ParseMode         string      `json:"parse_mode,omitempty"`         // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Width             int64       `json:"width,omitempty"`              // Optional. Video width
	Height            int64       `json:"height,omitempty"`             // Optional. Video height
	Duration          int64       `json:"duration,omitempty"`           // Optional. Video duration
	SupportsStreaming bool        `json:"supports_streaming,omitempty"` // Optional. Pass True, if the uploaded video is suitable for streaming
}
