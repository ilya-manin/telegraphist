package telegram

import (
	"os"
)

type SetChatPhotoParams struct {
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  *os.File    `json:"-"`       // New chat photo, uploaded using multipart/form-data
}

/*
Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
*/
func (c *Client) SetChatPhoto(params *SetChatPhotoParams) (bool, error) {
	var err error
	var result bool
	files := make(files)

	files.add("photo", params.Photo)

	err = c.performRequest("setChatPhoto", params, &files, &result)

	return result, err
}
