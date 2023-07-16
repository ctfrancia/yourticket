package main

import (
	"ctfrancia/yourticket/cmd/api/dto"
	"fmt"
	"net/http"
)

func (app *application) createFlightHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateFlightRequest
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Println(input)
	// TODO: validate input
	// err = app.models.Flights.Insert(input)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }
}

func (app *application) createAirport() {}

func (app *application) createAirline() {}

func (app *application) createCity() {}
