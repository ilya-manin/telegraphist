package telegram

/*
Voice represents a voice note.
*/
type Voice struct {
	FileID   string `json:"file_id"`             // Unique identifier for this file
	Duration int    `json:"duration"`            // Duration of the audio in seconds as defined by sender
	MimeType string `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize int    `json:"file_size,omitempty"` // Optional. File size
}
