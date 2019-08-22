package telegram

type SendChatActionParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string      `json:"action"`  // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
}

/*
Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
*/
func (c *Client) SendChatAction(params *SendChatActionParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("sendChatAction", params, nil, &result)

	return result, err
}
