package model

import "errors"

var (
	// ErrFlightAlreadyExists is an error that is returned when a flight is not found
	ErrNoFlightFound = errors.New("no flight found")
	// ErrDiff is just showing an additional error for this block declaration
	ErrDiff = errors.New("different error related to flight")
)

// CreateFlightResponse is the DTO response from the 3rd party provider regarding the creation of a flight
// within their system
type CreateFlightResponse struct {
	FlightNumber string `json:"flightNumber"`
}
