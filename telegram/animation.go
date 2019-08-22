package telegram

/*
Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
*/
type Animation struct {
	FileID   string     `json:"file_id"`             // Unique file identifier
	Width    int        `json:"width"`               // Video width as defined by sender
	Height   int        `json:"height"`              // Video height as defined by sender
	Duration int        `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Animation thumbnail as defined by sender
	FileName string     `json:"file_name,omitempty"` // Optional. Original animation filename as defined by sender
	MimeType string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize int        `json:"file_size,omitempty"` // Optional. File size
}
