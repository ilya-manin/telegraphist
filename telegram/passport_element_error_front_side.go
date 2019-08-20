package telegram

/*
PassportElementErrorFrontSide represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
*/
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`    // Error source, must be front_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the front side of the document
	Message  string `json:"message"`   // Error message
}
