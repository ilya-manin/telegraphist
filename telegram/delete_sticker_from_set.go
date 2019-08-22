package telegram

type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"` // File identifier of the sticker
}

/*
Use this method to delete a sticker from a set created by the bot. Returns True on success.
*/
func (c *Client) DeleteStickerFromSet(params *DeleteStickerFromSetParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("deleteStickerFromSet", params, nil, &result)

	return result, err
}
