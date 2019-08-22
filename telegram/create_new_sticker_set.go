package telegram

type CreateNewStickerSetParams struct {
	UserID        int           `json:"user_id"`                  // User identifier of created sticker set owner
	Name          string        `json:"name"`                     // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
	Title         string        `json:"title"`                    // Sticker set title, 1-64 characters
	PngSticker    interface{}   `json:"png_sticker"`              // Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Emojis        string        `json:"emojis"`                   // One or more emoji corresponding to the sticker
	ContainsMasks bool          `json:"contains_masks,omitempty"` // Pass True, if a set of mask stickers should be created
	MaskPosition  *MaskPosition `json:"mask_position,omitempty"`  // A JSON-serialized object for position where the mask should be placed on faces
}

/*
Use this method to create new sticker set owned by a user. The bot will be able to edit the created sticker set. Returns True on success.
*/
func (c *Client) CreateNewStickerSet(params *CreateNewStickerSetParams) (bool, error) {
	var err error
	var result bool
	files := make(files)

	err = handleFileToUpload(&params.PngSticker, &files)
	if err != nil {
		return result, err
	}

	err = c.performRequest("createNewStickerSet", params, &files, &result)

	return result, err
}
