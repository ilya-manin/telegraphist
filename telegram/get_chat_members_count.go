package telegram

type GetChatMembersCountParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

/*
Use this method to get the number of members in a chat. Returns Int on success.
*/
func (c *Client) GetChatMembersCount(params *GetChatMembersCountParams) (int, error) {
	var err error
	var result int

	err = c.performRequest("getChatMembersCount", params, nil, &result)

	return result, err
}
