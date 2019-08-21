package telegram

/*
A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
*/
func (c *Client) GetMe() (*User, error) {
	var err error
	var result User

	err = c.performRequest("getMe", nil, nil, &result)

	return &result, err
}
