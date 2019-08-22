package telegram

/*
UserProfilePhotos represent a user's profile pictures.
*/
type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]*PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}
