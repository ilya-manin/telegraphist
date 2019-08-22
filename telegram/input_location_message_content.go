package telegram

/*
InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query.
*/
type InputLocationMessageContent struct {
	Latitude   float64 `json:"latitude"`              // Latitude of the location in degrees
	Longitude  float64 `json:"longitude"`             // Longitude of the location in degrees
	LivePeriod int     `json:"live_period,omitempty"` // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
}
