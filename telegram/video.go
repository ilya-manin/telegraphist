package telegram

/*
Video represents a video file.
*/
type Video struct {
	FileID   string     `json:"file_id"`             // Unique identifier for this file
	Width    int64      `json:"width"`               // Video width as defined by sender
	Height   int64      `json:"height"`              // Video height as defined by sender
	Duration int64      `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Video thumbnail
	MimeType string     `json:"mime_type,omitempty"` // Optional. Mime type of a file as defined by sender
	FileSize int64      `json:"file_size,omitempty"` // Optional. File size
}
