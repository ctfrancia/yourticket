package mock

import (
	"ctfrancia/yourticket/cmd/api/dto"
	"ctfrancia/yourticket/internal/model"
	"ctfrancia/yourticket/internal/repository"
	"errors"
	"time"
)

var mockFlight = dto.CreateFlightRequest{
	FlightNumber: "AA123",
	Departure:    time.Now().AddDate(0, 0, -1),
	Arrival:      time.Now().AddDate(0, 0, 1),
	CreationDate: time.Now(),
	Price:        100.00,
	Airline:      "American Airlines",
}

// CreateFlightRequest is mocking a create flight request
func (r *repository.Repository) CreateFlight(f dto.CreateFlightRequest) (model.CreateFlightResponse, error) {
	switch f.FlightNumber {
	case "AA123":
		return model.CreateFlightResponse{FlightNumber: "AA123"}, nil
	case "AA456":
		return model.CreateFlightResponse{}, errors.New("flight already exists")
	default:
		// this default is to assume that there is a centralized flight records
		return model.CreateFlightResponse{}, errors.New("Uknown flight")
	}
	return model.CreateFlightResponse{FlightNumber: "AA123"}, nil
}
