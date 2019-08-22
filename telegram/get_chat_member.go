package telegram

type GetChatMemberParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserID int         `json:"user_id"` // Unique identifier of the target user
}

/*
Use this method to get information about a member of a chat. Returns a ChatMember object on success.
*/
func (c *Client) GetChatMember(params *GetChatMemberParams) (ChatMember, error) {
	var err error
	var result ChatMember

	err = c.performRequest("getChatMember", params, nil, &result)

	return result, err
}
