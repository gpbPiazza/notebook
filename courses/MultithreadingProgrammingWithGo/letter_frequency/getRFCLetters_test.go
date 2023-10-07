package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func setupMockServer(code int, response json.RawMessage) *httptest.Server {
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		})

	mockServer := httptest.NewServer(handler)
	return mockServer
}
