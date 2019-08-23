package telegram

type SendVenueParams struct {
	ChatID              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude            float64     `json:"latitude"`                       // Latitude of the venue
	Longitude           float64     `json:"longitude"`                      // Longitude of the venue
	Title               string      `json:"title"`                          // Name of the venue
	Address             string      `json:"address"`                        // Address of the venue
	FoursquareID        string      `json:"foursquare_id,omitempty"`        // Foursquare identifier of the venue
	FoursquareType      string      `json:"foursquare_type,omitempty"`      // Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	DisableNotification bool        `json:"disable_notification,omitempty"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`  // If the message is a reply, ID of the original message
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

/*
Use this method to send information about a venue. On success, the sent Message is returned.
*/
func (c *Client) SendVenue(params *SendVenueParams) (Message, error) {
	var err error
	var result Message

	err = c.performRequest("sendVenue", params, nil, &result)

	return result, err
}
