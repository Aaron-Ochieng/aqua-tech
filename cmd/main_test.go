package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterStaticFiles(t *testing.T) {
	mux := Router()
	type Tests struct {
		req    *http.Request
		status int
	}
	tests := []Tests{
		{
			httptest.NewRequest(http.MethodGet, "/static/test.css", nil), http.StatusNotFound,
		},
	}

	for _, test := range tests {

		w := httptest.NewRecorder()
		mux.ServeHTTP(w, test.req)
		if w.Result().StatusCode != test.status {
			t.Errorf("Expected status %v, got %v", w.Result().StatusCode, test.status)
		}
	}

}

func TestSensorDataEndpointValidPost(t *testing.T) {
	mux := Router()

	data := Data{
		Temp:           25.5,
		Humidity:       60.2,
		UltraSonicData: 123.45,
	}

	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/data", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Result().StatusCode)
	}
}

func TestSensorDataEndpointInvalidMethod(t *testing.T) {
	mux := Router()

	req := httptest.NewRequest(http.MethodGet, "/data", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status Method Not Allowed, got %v", w.Result().StatusCode)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["error"] != "Only POST method is allowed" {
		t.Errorf("Expected error message, got %v", response["error"])
	}
}
