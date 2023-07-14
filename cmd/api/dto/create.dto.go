package dto

import "time"

// CreateFlightRequest is the DTO for creating a flight request
type CreateFlightRequest struct {
	FlightNumber string    `json:"flightNumber"`
	Departure    string    `json:"departure"`
	Arrival      string    `json:"arrival"`
	Date         time.Time `json:"date"`
	Price        float32   `json:"price"`
	Airline      string    `json:"airline"`
}
