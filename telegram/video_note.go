package telegram

/*
VideoNote represents a video message (available in Telegram apps as of v.4.0).
*/
type VideoNote struct {
	FileID   string     `json:"file_id"`             // Unique identifier for this file
	Length   int64      `json:"length"`              // Video width and height (diameter of the video message) as defined by sender
	Duration int64      `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Video thumbnail
	FileSize int64      `json:"file_size,omitempty"` // Optional. File size
}