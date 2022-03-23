package utils

import (
	"net/http"
	"net/http/httptest"
)

// ExecuteRequest issues an http request against a test http recorder
func ExecuteRequest(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
