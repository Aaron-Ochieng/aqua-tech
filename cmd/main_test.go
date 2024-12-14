package main

import (
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
