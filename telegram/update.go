package telegram

/*
Update represents an incoming update.At most one of the optional parameters can be present in any given update.
*/
type Update struct {
	UpdateID           int                 `json:"update_id"`                      // The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message            *Message            `json:"message,omitempty"`              // Optional. New incoming message of any kind — text, photo, sticker, etc.
	EditedMessage      *Message            `json:"edited_message,omitempty"`       // Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *Message            `json:"channel_post,omitempty"`         // Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`  // Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`         // Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`       // Optional. New incoming callback query
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`       // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`   // Optional. New incoming pre-checkout query. Contains full information about checkout
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
}
