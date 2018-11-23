package main

import (
	"log"
	"net/http"
	"os"

	"./handler"
)

func main() {

	port := os.Getenv("PORT") //OS
	if port == "" {
		port = "8080"
	}

	myHandler := handler.NewHandler()

	server := &http.Server{
		Addr:    ":8080",
		Handler: myHandler,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
