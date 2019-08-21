package telegram

/*
SendPhotoParams is a representation of request parameters for SendPhoto
*/
type SendPhotoParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo               interface{} `json:"photo"`                          // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. More info on Sending Files »
	Caption             string      `json:"caption,omitempty"`              // Photo caption (may also be used when resending photos by file_id), 0-1024 characters
	ParseMode           string      `json:"parse_mode,omitempty"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send photos. On success, the sent Message is returned.
*/
func (c *Client) SendPhoto(params *SendPhotoParams) (*Message, error) {
	var err error
	var result Message
	files := make(files)

	err = handleFileToUpload(&params.Photo, &files)
	if err != nil {
		return &result, err
	}

	err = c.performRequest("sendPhoto", params, &files, &result)

	return &result, err
}
