package telegram

/*
PassportElementErrorTranslationFile represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
*/
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`    // Error source, must be translation_file
	Type     string `json:"type"`      // Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}
