package telegram

type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`  // File identifier of the sticker
	Position int    `json:"position"` // New sticker position in the set, zero-based
}

/*
Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.
*/
func (c *Client) SetStickerPositionInSet(params *SetStickerPositionInSetParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("setStickerPositionInSet", params, nil, &result)

	return result, err
}
