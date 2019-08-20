package telegram

/*
File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
*/
type File struct {
	FileID   string `json:"file_id"`             // Unique identifier for this file
	FileSize int64  `json:"file_size,omitempty"` // Optional. File size, if known
	FilePath string `json:"file_path,omitempty"` // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}
