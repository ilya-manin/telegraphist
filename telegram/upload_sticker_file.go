package telegram

import (
	"os"
)

type UploadStickerFileParams struct {
	UserID     int      `json:"user_id"` // User identifier of sticker file owner
	PngSticker *os.File `json:"-"`       // Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »
}

/*
Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
*/
func (c *Client) UploadStickerFile(params *UploadStickerFileParams) (File, error) {
	var err error
	var result File
	files := make(files)

	files.add("png_sticker", params.PngSticker)

	err = c.performRequest("uploadStickerFile", params, &files, &result)

	return result, err
}
