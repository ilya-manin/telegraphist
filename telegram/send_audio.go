package telegram

type SendAudioParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Audio               interface{} `json:"audio"`                          // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Caption             string      `json:"caption,omitempty"`              // Audio caption, 0-1024 characters
	ParseMode           string      `json:"parse_mode,omitempty"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration            int         `json:"duration,omitempty"`             // Duration of the audio in seconds
	Performer           string      `json:"performer,omitempty"`            // Performer
	Title               string      `json:"title,omitempty"`                // Track name
	Thumb               interface{} `json:"thumb,omitempty"`                // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
*/
func (c *Client) SendAudio(params *SendAudioParams) (Message, error) {
	var err error
	var result Message
	files := make(files)

	err = handleFileToUpload(&params.Audio, &files)
	if err != nil {
		return result, err
	}

	if params.Thumb != nil {
		err = handleFileToUpload(&params.Thumb, &files)
		if err != nil {
			return result, err
		}
	}

	err = c.performRequest("sendAudio", params, &files, &result)

	return result, err
}
