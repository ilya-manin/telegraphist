package telegram

/*
ChatPhoto represents a chat photo.
*/
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"` // Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileID   string `json:"big_file_id"`   // Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
}
