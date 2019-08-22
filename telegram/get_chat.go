package telegram

type GetChatParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

/*
Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
*/
func (c *Client) GetChat(params *GetChatParams) (Chat, error) {
	var err error
	var result Chat

	err = c.performRequest("getChat", params, nil, &result)

	return result, err
}
