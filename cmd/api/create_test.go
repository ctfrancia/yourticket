package main

import (
	"testing"
)

func Test_CreateFlight(t *testing.T) {
	// TODO: Implement
	// here we would create a test server that we would import from internal/mock
	// create a test that would test the create flight handler
	// test the create flight handler with a bad request
	// test the create flight handler with a server error
	// test the create flight handler with a successful request
	// test the create flight handler with a flight already exists
	// for end-to-end testing
	/*
		app := newTestApplication(t)
		ts := newTestServer(t, app.routes())
		defer ts.Close()

		tests := []struct {
			name     string
			urlPath  string
			body     dto.CreateFlightRequest
			codeWant int
			wantBody []byte
		}{
			{"Valid Flight", "/api/v1/flights", dto.CreateFlightRequest{ valid body here }, http.StatusOK, []byte(`{ "flight": "flightNumber":"AA123"}`)},
			{"Duplicate Flight", "/api/v1/flights", dto.CreateFlightRequest{}, http.StatusConflict, nil},
			{"Invalid Flight", "/api/v1/flights", dto.CreateFlightRequest{}, http.StatusBadRequest, nil},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				code, _, body := ts.get(t, tt.urlPath)
				if code != tt.codeWant {
					t.Errorf("wanted: %d; got: %d", tt.wantBody, code)
				}

				if !bytes.Contains(body, tt.wantBody) {
					t.Errorf("Want body: %q, received body: %q", tt.wantBody, body)
				}
			})
		}

	*/
}
