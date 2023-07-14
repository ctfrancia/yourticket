package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundRequest)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "api/v1/flights/airport/:aCode", app.flightsInAirportHandler)
	router.HandlerFunc(http.MethodGet, "api/v1/flights/city/:cCode", app.flightsByCityHandler)

	router.HandlerFunc(http.MethodPost, "api/v1/flights", app.createFlightHandler)
	return router
}
