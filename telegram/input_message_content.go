package telegram

/*
InputMessageContent represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 4 types:
*/
type InputMessageContent struct {
	MessageText           string `json:"message_text"`                       // Text of the message to be sent, 1-4096 characters
	ParseMode             string `json:"parse_mode,omitempty"`               // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in the sent message
}
