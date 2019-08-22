package telegram

/*
Audio represents an audio file to be treated as music by the Telegram clients.
*/
type Audio struct {
	FileID    string     `json:"file_id"`             // Unique identifier for this file
	Duration  int        `json:"duration"`            // Duration of the audio in seconds as defined by sender
	Performer string     `json:"performer,omitempty"` // Optional. Performer of the audio as defined by sender or by audio tags
	Title     string     `json:"title,omitempty"`     // Optional. Title of the audio as defined by sender or by audio tags
	MimeType  string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize  int        `json:"file_size,omitempty"` // Optional. File size
	Thumb     *PhotoSize `json:"thumb,omitempty"`     // Optional. Thumbnail of the album cover to which the music file belongs
}
