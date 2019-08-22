package telegram

type ExportChatInviteLinkParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

/*
Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.
*/
func (c *Client) ExportChatInviteLink(params *ExportChatInviteLinkParams) (string, error) {
	var err error
	var result string

	err = c.performRequest("exportChatInviteLink", params, nil, &result)

	return result, err
}
