package telegram

/*
InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
*/
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`                            // Type of the result, must be sticker
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	StickerFileID       string                `json:"sticker_file_id"`                 // A valid file identifier of the sticker
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the sticker
}
