package telegram

/*
Venue represents a venue.
*/
type Venue struct {
	Location       *Location `json:"location"`                  // Venue location
	Title          string    `json:"title"`                     // Name of the venue
	Address        string    `json:"address"`                   // Address of the venue
	FoursquareID   string    `json:"foursquare_id,omitempty"`   // Optional. Foursquare identifier of the venue
	FoursquareType string    `json:"foursquare_type,omitempty"` // Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
}
