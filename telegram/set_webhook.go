package telegram

type SetWebhookParams struct {
	URL            string      `json:"url"`                       // HTTPS url to send updates to. Use an empty string to remove webhook integration
	Certificate    interface{} `json:"certificate,omitempty"`     // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	MaxConnections int         `json:"max_connections,omitempty"` // Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput.
	AllowedUpdates []string    `json:"allowed_updates,omitempty"` // List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
}

/*
Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
*/
func (c *Client) SetWebhook(params *SetWebhookParams) (bool, error) {
	var err error
	var result bool
	files := make(files)

	if params.Certificate != nil {
		err = handleFileToUpload(&params.Certificate, &files)
		if err != nil {
			return result, err
		}
	}

	err = c.performRequest("setWebhook", params, &files, &result)

	return result, err
}
