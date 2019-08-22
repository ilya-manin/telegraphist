package telegram

type PinChatMessageParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID           int         `json:"message_id"`                     // IDentifier of a message to pin
	DisableNotification bool        `json:"disable_notification,omitempty"` // Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
}

/*
Use this method to pin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
*/
func (c *Client) PinChatMessage(params *PinChatMessageParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("pinChatMessage", params, nil, &result)

	return result, err
}
