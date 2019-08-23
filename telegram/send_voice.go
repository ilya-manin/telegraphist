package telegram

type SendVoiceParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Voice               interface{} `json:"voice"`                          // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Caption             string      `json:"caption,omitempty"`              // Voice message caption, 0-1024 characters
	ParseMode           string      `json:"parse_mode,omitempty"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration            int         `json:"duration,omitempty"`             // Duration of the voice message in seconds
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
*/
func (c *Client) SendVoice(params *SendVoiceParams) (Message, error) {
	var err error
	var result Message
	files := make(files)

	err = handleFileToUpload(&params.Voice, &files)
	if err != nil {
		return result, err
	}

	err = c.performRequest("sendVoice", params, &files, &result)

	return result, err
}
