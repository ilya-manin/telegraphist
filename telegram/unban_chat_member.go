package telegram

type UnbanChatMemberParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
	UserID int         `json:"user_id"` // Unique identifier of the target user
}

/*
Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.
*/
func (c *Client) UnbanChatMember(params *UnbanChatMemberParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("unbanChatMember", params, nil, &result)

	return result, err
}
