package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Arthur1/sendgrid-events-to-mackerel/internal/server"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	h := server.NewHTTPHandler()
	log.Fatal(http.ListenAndServe(":"+port, h))
}
