package telegram

/*
Poll contains information about a poll.
*/
type Poll struct {
	ID       string        `json:"id"`        // Unique poll identifier
	Question string        `json:"question"`  // Poll question, 1-255 characters
	Options  []*PollOption `json:"options"`   // List of poll options
	IsClosed bool          `json:"is_closed"` // True, if the poll is closed
}
