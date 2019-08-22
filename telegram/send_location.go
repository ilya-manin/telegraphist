package telegram

type SendLocationParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude            float64     `json:"latitude"`                       // Latitude of the location
	Longitude           float64     `json:"longitude"`                      // Longitude of the location
	LivePeriod          int         `json:"live_period,omitempty"`          // Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send point on the map. On success, the sent Message is returned.
*/
func (c *Client) SendLocation(params *SendLocationParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("sendLocation", params, nil, &result)

	return result, err
}
