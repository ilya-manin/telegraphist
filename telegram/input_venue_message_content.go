package telegram

/*
InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query.
*/
type InputVenueMessageContent struct {
	Latitude       float64 `json:"latitude"`                  // Latitude of the venue in degrees
	Longitude      float64 `json:"longitude"`                 // Longitude of the venue in degrees
	Title          string  `json:"title"`                     // Name of the venue
	Address        string  `json:"address"`                   // Address of the venue
	FoursquareID   string  `json:"foursquare_id,omitempty"`   // Optional. Foursquare identifier of the venue, if known
	FoursquareType string  `json:"foursquare_type,omitempty"` // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
}
