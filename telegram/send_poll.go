package telegram

type SendPollParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername). A native poll can't be sent to a private chat.
	Question            string      `json:"question"`                       // Poll question, 1-255 characters
	Options             []string    `json:"options"`                        // List of answer options, 2-10 strings 1-100 characters each
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send a native poll. A native poll can't be sent to a private chat. On success, the sent Message is returned.
*/
func (c *Client) SendPoll(params *SendPollParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("sendPoll", params, nil, &result)
	return result, err
}
