package repository

import (
	"bytes"
	"ctfrancia/yourticket/cmd/api/dto"
	"ctfrancia/yourticket/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type repository interface {
	CreateFlight() (model.CreateFlightResponse, error)
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
		flightsBaseUrl: "https://34e7a5ef-5e16-4dd8-88ad-d2f49ddc33be.mock.pstmn.io",
	}
}

// CreateFlight calls a 3rd party url to create a new flight at 3rd party or db or whatever
//
func (r *Repository) CreateFlight(f dto.CreateFlightRequest) (model.CreateFlightResponse, error) {
	var response model.CreateFlightResponse
	postBody, err := json.Marshal(f)
	if err != nil {
		return model.CreateFlightResponse{}, err
	}

	responseBody := bytes.NewBuffer(postBody)
	url := fmt.Sprintf("%s/%s/%s/create", r.flightsBaseUrl, v1, tenent)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return model.CreateFlightResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &response)
	return model.CreateFlightResponse{
		FlightNumber: f.FlightNumber,
	}, nil
}
