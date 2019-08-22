package telegram

type GetChatAdministratorsParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

/*
Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
*/
func (c *Client) GetChatAdministrators(params *GetChatAdministratorsParams) ([]ChatMember, error) {
	var err error
	var result []ChatMember

	err = c.performRequest("getChatAdministrators", params, nil, &result)

	return result, err
}
