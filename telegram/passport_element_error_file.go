package telegram

/*
PassportElementErrorFile represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
*/
type PassportElementErrorFile struct {
	Source   string `json:"source"`    // Error source, must be file
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}
