package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFound(t *testing.T) {
	handler := NewHandler()

	req, _ := http.NewRequest("GET", "/doesnotexist", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if res.Code != 404 {
		t.Fatalf("Wrong status code returned: %d", res.Code)
	}
}
