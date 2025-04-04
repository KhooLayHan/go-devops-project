package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KhooLayHan/go-devops-project/internal/app"
	"github.com/KhooLayHan/go-devops-project/internal/version"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", app.Handler)
	http.HandleFunc("/version", version.Handler)

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}