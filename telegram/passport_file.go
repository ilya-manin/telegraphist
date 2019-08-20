package telegram

/*
PassportFile represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
*/
type PassportFile struct {
	FileID   string `json:"file_id"`   // Unique identifier for this file
	FileSize int64  `json:"file_size"` // File size
	FileDate int64  `json:"file_date"` // Unix time when the file was uploaded
}
