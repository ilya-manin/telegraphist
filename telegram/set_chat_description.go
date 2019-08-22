package telegram

type SetChatDescriptionParams struct {
	ChatID      interface{} `json:"chat_id"`     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description string      `json:"description"` // New chat description, 0-255 characters
}

/*
Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
*/
func (c *Client) SetChatDescription(params *SetChatDescriptionParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("setChatDescription", params, nil, &result)

	return result, err
}
