package telegram

type AddStickerToSetParams struct {
	UserID       int           `json:"user_id"`                 // User identifier of sticker set owner
	Name         string        `json:"name"`                    // Sticker set name
	PngSticker   interface{}   `json:"png_sticker"`             // Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Emojis       string        `json:"emojis"`                  // One or more emoji corresponding to the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // A JSON-serialized object for position where the mask should be placed on faces
}

/*
Use this method to add a new sticker to a set created by the bot. Returns True on success.
*/
func (c *Client) AddStickerToSet(params *AddStickerToSetParams) (bool, error) {
	var err error
	var result bool
	files := make(files)

	err = handleFileToUpload(&params.PngSticker, &files)
	if err != nil {
		return result, err
	}

	err = c.performRequest("addStickerToSet", params, &files, &result)

	return result, err
}
