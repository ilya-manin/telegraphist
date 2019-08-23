package telegram

type GetUserProfilePhotosParams struct {
	UserID int `json:"user_id"`          // Unique identifier of the target user
	Offset int `json:"offset,omitempty"` // Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int `json:"limit,omitempty"`  // Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
}

/*
Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
*/
func (c *Client) GetUserProfilePhotos(params *GetUserProfilePhotosParams) (UserProfilePhotos, error) {
	var err error
	var result UserProfilePhotos

	err = c.performRequest("getUserProfilePhotos", params, nil, &result)

	return result, err
}
