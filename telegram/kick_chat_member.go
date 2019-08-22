package telegram

type KickChatMemberParams struct {
	ChatID    interface{} `json:"chat_id"`              // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID    int         `json:"user_id"`              // Unique identifier of the target user
	UntilDate int         `json:"until_date,omitempty"` // Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
}

/*
Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
*/
func (c *Client) KickChatMember(params *KickChatMemberParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("kickChatMember", params, nil, &result)

	return result, err
}
