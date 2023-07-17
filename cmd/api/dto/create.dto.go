package dto

import "time"

// CreateFlightRequest is the DTO for creating a flight request
type CreateFlightRequest struct {
	FlightNumber string    `json:"flightNumber"`
	Departure    time.Time `json:"departure"`
	Arrival      time.Time `json:"arrival"`
	CreationDate time.Time `json:"creationDate"`
	Price        float32   `json:"price"`
	Airline      string    `json:"airline"`
}
