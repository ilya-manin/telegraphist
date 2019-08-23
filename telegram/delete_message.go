package telegram

type DeleteMessageParams struct {
	ChatID    interface{} `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID int         `json:"message_id"` // IDentifier of the message to delete
}

/*
Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
*/
func (c *Client) DeleteMessage(params *DeleteMessageParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("deleteMessage", params, nil, &result)

	return result, err
}
