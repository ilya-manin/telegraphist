package telegram

/*
Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.
*/
func (c *Client) DeleteWebhook() (bool, error) {
	var err error
	var result bool

	err = c.performRequest("deleteWebhook", nil, nil, &result)

	return result, err
}
