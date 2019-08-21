package telegram

/*
Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
*/
func (c *Client) GetWebhookInfo() (WebhookInfo, error) {
	var err error
	var result WebhookInfo

	err = c.performRequest("getWebhookInfo", nil, nil, &result)

	return result, err
}
