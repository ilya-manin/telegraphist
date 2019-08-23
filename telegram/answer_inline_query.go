package telegram

type AnswerInlineQueryParams struct {
	InlineQueryID     string               `json:"inline_query_id"`               // Unique identifier for the answered query
	Results           *[]InlineQueryResult `json:"results"`                       // A JSON-serialized array of results for the inline query
	CacheTime         int                  `json:"cache_time,omitempty"`          // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal        bool                 `json:"is_personal,omitempty"`         // Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	NextOffset        string               `json:"next_offset,omitempty"`         // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
	SwitchPmText      string               `json:"switch_pm_text,omitempty"`      // If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmParameter string               `json:"switch_pm_parameter,omitempty"` // Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

/*
Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
*/
func (c *Client) AnswerInlineQuery(params *AnswerInlineQueryParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("answerInlineQuery", params, nil, &result)

	return result, err
}
