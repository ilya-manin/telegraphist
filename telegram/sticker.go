package telegram

/*
Sticker represents a sticker.
*/
type Sticker struct {
	FileID       string        `json:"file_id"`                 // Unique identifier for this file
	Width        int64         `json:"width"`                   // Sticker width
	Height       int64         `json:"height"`                  // Sticker height
	IsAnimated   bool          `json:"is_animated"`             // True, if the sticker is animated
	Thumb        *PhotoSize    `json:"thumb,omitempty"`         // Optional. Sticker thumbnail in the .webp or .jpg format
	Emoji        string        `json:"emoji,omitempty"`         // Optional. Emoji associated with the sticker
	SetName      string        `json:"set_name,omitempty"`      // Optional. Name of the sticker set to which the sticker belongs
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional. For mask stickers, the position where the mask should be placed
	FileSize     int64         `json:"file_size,omitempty"`     // Optional. File size
}
