package telegram

/*
ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
*/
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`                   // The unique identifier for the result that was chosen
	From            *User     `json:"from"`                        // The user that chose the result
	Location        *Location `json:"location,omitempty"`          // Optional. Sender location, only for bots that require user location
	InlineMessageID string    `json:"inline_message_id,omitempty"` // Optional. IDentifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query"`                       // The query that was used to obtain the result
}
