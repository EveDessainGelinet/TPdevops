package handler

import (
	"net/http"
)

func NewHandler() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("/hello", Hello)
	handler.HandleFunc("/", NotFound)

	return handler
}
