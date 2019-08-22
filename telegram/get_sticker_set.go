package telegram

type GetStickerSetParams struct {
	Name string `json:"name"` // Name of the sticker set
}

/*
Use this method to get a sticker set. On success, a StickerSet object is returned.
*/
func (c *Client) GetStickerSet(params *GetStickerSetParams) (StickerSet, error) {
	var err error
	var result StickerSet

	err = c.performRequest("getStickerSet", params, nil, &result)

	return result, err
}
