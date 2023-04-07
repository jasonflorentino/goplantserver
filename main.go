package main

import (
	"fmt"
	"net/http"

	"github.com/jasonflorentino/goplantserver/handlers"
	"github.com/jasonflorentino/goplantserver/logger"
	"github.com/jasonflorentino/goplantserver/middlewares"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/", middlewares.UseGlobalMiddlewares(helloHandler))
	router.Handle("/about", middlewares.UseGlobalMiddlewares(aboutHandler))

	router.Handle("/plants", middlewares.ProtectRoute(handlers.PlantsHandler))

	logger.Logger.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Logger.Fatal("Error starting server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}
