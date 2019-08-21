package telegram

type GetUpdatesParams struct {
	Offset         int      `json:"offset,omitempty"`          // IDentifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Limit          int      `json:"limit,omitempty"`           // Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Timeout        int      `json:"timeout,omitempty"`         // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates,omitempty"` // List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

/*
Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.
*/
func (c *Client) GetUpdates(params *GetUpdatesParams) ([]Update, error) {
	var err error
	var result []Update

	err = c.performRequest("getUpdates", params, nil, &result)

	return result, err
}
