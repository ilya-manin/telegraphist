package telegram

type LeaveChatParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

/*
Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
*/
func (c *Client) LeaveChat(params *LeaveChatParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("leaveChat", params, nil, &result)

	return result, err
}
