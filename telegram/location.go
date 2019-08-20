package telegram

/*
Location represents a point on the map.
*/
type Location struct {
	Longitude float64 `json:"longitude"` // Longitude as defined by sender
	Latitude  float64 `json:"latitude"`  // Latitude as defined by sender
}
