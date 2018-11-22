package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	myHandler := NewHandler()

	req, _ := http.NewRequest("GET", "/hello", nil)
	res := httptest.NewRecorder()

	myHandler.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Fatalf("Wrong status code returned: %d", res.Code)
	}
}
