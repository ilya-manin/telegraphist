package telegram

type SendVideoNoteParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	VideoNote           interface{} `json:"video_note"`                     // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration            int         `json:"duration,omitempty"`             // Duration of sent video in seconds
	Length              int         `json:"length,omitempty"`               // Video width and height, i.e. diameter of the video message
	Thumb               interface{} `json:"thumb,omitempty"`                // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
*/
func (c *Client) SendVideoNote(params *SendVideoNoteParams) (Message, error) {
	var err error
	var result Message
	files := make(files)

	err = handleFileToUpload(&params.VideoNote, &files)
	if err != nil {
		return result, err
	}

	if params.Thumb != nil {
		err = handleFileToUpload(&params.Thumb, &files)
		if err != nil {
			return result, err
		}
	}

	err = c.performRequest("sendVideoNote", params, &files, &result)

	return result, err
}
