package telegram

/*
Document represents a general file (as opposed to photos, voice messages and audio files).
*/
type Document struct {
	FileID   string     `json:"file_id"`             // Unique file identifier
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Document thumbnail as defined by sender
	FileName string     `json:"file_name,omitempty"` // Optional. Original filename as defined by sender
	MimeType string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize int        `json:"file_size,omitempty"` // Optional. File size
}
