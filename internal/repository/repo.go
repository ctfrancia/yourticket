package repository

import (
	"ctfrancia/yourticket/cmd/api/dto"
	"fmt"
)

type repository interface {
	// FindByCode(string) (*entity.Coupon, error)
	// Save(entity.Coupon) error
	// Close() error
	CreateFlight() error
}

const (
	v1     = "v1"
	tenent = "flights"
)

// Repository is the struct that holds the repository
type Repository struct {
	flightsBaseUrl string
}

// New creates a new repository instance
func New() *Repository {
	return &Repository{
		flightsBaseUrl: "http://example-of-3rd-party-url.com/api/v1/",
	}
}

// CreateFlight calls a 3rd party url to create a new flight at 3rd party or db or whatever
func (r *Repository) CreateFlight(f dto.CreateFlightRequest) error {
	url, _ := fmt.Printf("%s/%s/%s/create", r.flightsBaseUrl, v1, tenent)
	fmt.Println("url", url)
	return nil
}
