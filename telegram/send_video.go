package telegram

type SendVideoParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Video               interface{} `json:"video"`                          // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
	Duration            int         `json:"duration,omitempty"`             // Duration of sent video in seconds
	Width               int         `json:"width,omitempty"`                // Video width
	Height              int         `json:"height,omitempty"`               // Video height
	Thumb               interface{} `json:"thumb,omitempty"`                // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption             string      `json:"caption,omitempty"`              // Video caption (may also be used when resending videos by file_id), 0-1024 characters
	ParseMode           string      `json:"parse_mode,omitempty"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	SupportsStreaming   bool        `json:"supports_streaming,omitempty"`   // Pass True, if the uploaded video is suitable for streaming
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
*/
func (c *Client) SendVideo(params *SendVideoParams) (Message, error) {
	var err error
	var result Message
	files := make(files)

	err = handleFileToUpload(&params.Video, &files)
	if err != nil {
		return result, err
	}

	if params.Thumb != nil {
		err = handleFileToUpload(&params.Thumb, &files)
		if err != nil {
			return result, err
		}
	}

	err = c.performRequest("sendVideo", params, &files, &result)

	return result, err
}