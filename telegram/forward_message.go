package telegram

type ForwardMessageParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatID          interface{} `json:"from_chat_id"`                   // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	MessageID           int         `json:"message_id"`                     // Message identifier in the chat specified in from_chat_id
}

/*
Use this method to forward messages of any kind. On success, the sent Message is returned.
*/
func (c *Client) ForwardMessage(params *ForwardMessageParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("forwardMessage", params, nil, &result)

	return result, err
}
