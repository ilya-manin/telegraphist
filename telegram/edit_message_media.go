package telegram

import (
	"fmt"
)

type EditMessageMediaParams struct {
	ChatID          interface{}           `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       int                   `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. IDentifier of the message to edit
	InlineMessageID string                `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. IDentifier of the inline message
	Media           interface{}           `json:"media"`                       // A JSON-serialized object for a new media content of the message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // A JSON-serialized object for a new inline keyboard.
}

/*
Use this method to edit animation, audio, document, photo, or video messages. If a message is a part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (c *Client) EditMessageMedia(params *EditMessageMediaParams) (Message, error) {
	var err error
	var result Message
	files := make(files)

	if animation, ok := params.Media.(*InputMediaAnimation); ok {
		err = handleFileToUpload(&animation.Media, &files)
		if err != nil {
			return result, err
		}
		if animation.Thumb != nil {
			err = handleFileToUpload(&animation.Thumb, &files)
			if err != nil {
				return result, err
			}
		}
	} else if audio, ok := params.Media.(*InputMediaAudio); ok {
		err = handleFileToUpload(&audio.Media, &files)
		if err != nil {
			return result, err
		}
		if audio.Thumb != nil {
			err = handleFileToUpload(&audio.Thumb, &files)
			if err != nil {
				return result, err
			}
		}
	} else if document, ok := params.Media.(*InputMediaDocument); ok {
		err = handleFileToUpload(&document.Media, &files)
		if err != nil {
			return result, err
		}
		if document.Thumb != nil {
			err = handleFileToUpload(&document.Thumb, &files)
			if err != nil {
				return result, err
			}
		}
	} else if photo, ok := params.Media.(*InputMediaPhoto); ok {
		err = handleFileToUpload(&photo.Media, &files)
		if err != nil {
			return result, err
		}
	} else if video, ok := params.Media.(*InputMediaVideo); ok {
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
		return result, fmt.Errorf("Unknow media type, should be InputMediaAnimation, InputMediaAudio, InputMediaDocument, InputMediaPhoto or InputMediaVideo")
	}

	err = c.performRequest("editMessageMedia", params, &files, &result)

	return result, err
}
