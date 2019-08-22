package telegram

/*
PollOption contains information about one answer option in a poll.
*/
type PollOption struct {
	Text       string `json:"text"`        // Option text, 1-100 characters
	VoterCount int    `json:"voter_count"` // Number of users that voted for this option
}
