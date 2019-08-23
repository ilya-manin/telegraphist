package telegram

type SetPassportDataErrorsParams struct {
	UserID int                     `json:"user_id"` // User identifier
	Errors *[]PassportElementError `json:"errors"`  // A JSON-serialized array describing the errors
}

/*
Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
*/
func (c *Client) SetPassportDataErrors(params *SetPassportDataErrorsParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("setPassportDataErrors", params, nil, &result)

	return result, err
}
