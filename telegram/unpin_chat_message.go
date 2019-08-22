package telegram

type UnpinChatMessageParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

/*
Use this method to unpin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
*/
func (c *Client) UnpinChatMessage(params *UnpinChatMessageParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("unpinChatMessage", params, nil, &result)

	return result, err
}
