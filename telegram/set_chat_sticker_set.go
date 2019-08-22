package telegram

type SetChatStickerSetParams struct {
	ChatID         interface{} `json:"chat_id"`          // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string      `json:"sticker_set_name"` // Name of the sticker set to be set as the group sticker set
}

/*
Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
*/
func (c *Client) SetChatStickerSet(params *SetChatStickerSetParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("setChatStickerSet", params, nil, &result)

	return result, err
}
