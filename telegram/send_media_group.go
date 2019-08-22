package telegram

import (
	"fmt"
)

/*
SendMediaGroupParams is a representation of request parameters for SendMediaGroup
*/
type SendMediaGroupParams struct {
	ChatID              interface{}   `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Media               []interface{} `json:"media"`                          // A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
	DisableNotification bool          `json:"disable_notification,omitempty"` // Sends the messages silently. Users will receive a notification with no sound.
	ReplyToMessageID    int           `json:"reply_to_message_id,omitempty"`  // If the messages are a reply, ID of the original message
}

/*
Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
*/
func (c *Client) SendMediaGroup(params *SendMediaGroupParams) ([]Message, error) {
	var err error
	var result []Message
	files := make(files)

	for _, media := range params.Media {
		if photo, ok := media.(*InputMediaPhoto); ok {
			err = handleFileToUpload(&photo.Media, &files)
			if err != nil {
				return result, err
			}
		} else if video, ok := media.(*InputMediaVideo); ok {
			err = handleFileToUpload(&video.Media, &files)
			if err != nil {
				return result, err
			}
			if video.Thumb != nil {
				err = handleFileToUpload(&video.Thumb, &files)
				if err != nil {
					return result, err
				}
			}
		} else {
			return result, fmt.Errorf("Unknow media type, should be InputMediaPhoto or InputMediaVideo")
		}
	}

	err = c.performRequest("sendMediaGroup", params, &files, &result)

	return result, err
}
