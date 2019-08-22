package telegram

type RestrictChatMemberParams struct {
	ChatID      interface{}      `json:"chat_id"`              // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID      int              `json:"user_id"`              // Unique identifier of the target user
	Permissions *ChatPermissions `json:"permissions"`          // New user permissions
	UntilDate   int              `json:"until_date,omitempty"` // Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

/*
Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
*/
func (c *Client) RestrictChatMember(params *RestrictChatMemberParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("restrictChatMember", params, nil, &result)

	return result, err
}
