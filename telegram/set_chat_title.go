package telegram

type SetChatTitleParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string      `json:"title"`   // New chat title, 1-255 characters
}

/*
Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
*/
func (c *Client) SetChatTitle(params *SetChatTitleParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("setChatTitle", params, nil, &result)

	return result, err
}
