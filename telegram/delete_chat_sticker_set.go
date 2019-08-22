package telegram

type DeleteChatStickerSetParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

/*
Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
*/
func (c *Client) DeleteChatStickerSet(params *DeleteChatStickerSetParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("deleteChatStickerSet", params, nil, &result)

	return result, err
}
