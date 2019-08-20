package telegram

/*
InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
*/
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}
