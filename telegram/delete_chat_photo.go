package telegram

type DeleteChatPhotoParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

/*
Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
*/
func (c *Client) DeleteChatPhoto(params *DeleteChatPhotoParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("deleteChatPhoto", params, nil, &result)

	return result, err
}
