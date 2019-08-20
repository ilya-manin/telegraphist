package telegram

/*
PassportElementErrorUnspecified represents an issue in an unspecified place. The error is considered resolved when new data is added.
*/
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`       // Error source, must be unspecified
	Type        string `json:"type"`         // Type of element of the user's Telegram Passport which has the issue
	ElementHash string `json:"element_hash"` // Base64-encoded element hash
	Message     string `json:"message"`      // Error message
}
