package telegram

/*
Video represents a video file.
*/
type Video struct {
	FileID   string     `json:"file_id"`             // Unique identifier for this file
	Width    int        `json:"width"`               // Video width as defined by sender
	Height   int        `json:"height"`              // Video height as defined by sender
	Duration int        `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Video thumbnail
	MimeType string     `json:"mime_type,omitempty"` // Optional. Mime type of a file as defined by sender
	FileSize int        `json:"file_size,omitempty"` // Optional. File size
}
