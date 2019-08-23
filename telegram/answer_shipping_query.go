package telegram

type AnswerShippingQueryParams struct {
	ShippingQueryID string            `json:"shipping_query_id"`          // Unique identifier for the query to be answered
	Ok              bool              `json:"ok"`                         // Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions *[]ShippingOption `json:"shipping_options,omitempty"` // Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string            `json:"error_message,omitempty"`    // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
}

/*
If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
*/
func (c *Client) AnswerShippingQuery(params *AnswerShippingQueryParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("answerShippingQuery", params, nil, &result)

	return result, err
}
