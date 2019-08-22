package telegram

type SetChatPermissionsParams struct {
	ChatID      interface{}      `json:"chat_id"`     // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions *ChatPermissions `json:"permissions"` // New default chat permissions
}

/*
Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members admin rights. Returns True on success.
*/
func (c *Client) SetChatPermissions(params *SetChatPermissionsParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("setChatPermissions", params, nil, &result)

	return result, err
}
