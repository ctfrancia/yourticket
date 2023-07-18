package main

import (
	"ctfrancia/yourticket/cmd/api/dto"
	"fmt"
	"net/http"
)

func (app *application) createFlightHandler(w http.ResponseWriter, r *http.Request) {
	var cfr dto.CreateFlightRequest
	err := app.readJSON(w, r, &cfr)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	resp, err := app.repo.CreateFlight(cfr)
	if err != nil {
		// co
		log := fmt.Sprintf("error within repo: %s \n, %s", cfr.FlightNumber, err.Error())
		app.logger.Println(log)
		// resp := app.logger.(err).Errorf("unable to serve HTTP POST request for flight %s", cfr.FlightNumber)
		app.serverErrorResponse(w, r, err)
		return
	}
	ret := envelope{"flight": resp}
	err = app.writeJSON(w, http.StatusCreated, ret, nil)
}

// The following functions are not implemented yet
// this is on purpose to show some other APIs that are created
// for the various ways that we can create and think about future
// APIs
func (app *application) createAirport() {}
func (app *application) createAirline() {}
func (app *application) createCity()    {}
