package main

import (
	"fmt"
	"net/http"

	"github.com/jasonflorentino/goplantserver/handlers"
	"github.com/jasonflorentino/goplantserver/logger"
)

func main() {
	http.HandleFunc("/", helloHandler)      // register the original handler for the root path
	http.HandleFunc("/about", aboutHandler) // register a new handler for the "/about" path
	http.HandleFunc("/plants", handlers.PlantsHandler)

	logger.Logger.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Logger.Fatal("Error starting server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Println("Handling request:", r.URL.Path)
	fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Println("Handling request:", r.URL.Path)
	fmt.Fprintf(w, "This is the about page.")
}
