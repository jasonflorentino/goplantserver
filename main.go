package main

import (
	"net/http"

	"github.com/jasonflorentino/goplantserver/handlers"
	"github.com/jasonflorentino/goplantserver/logger"
	"github.com/jasonflorentino/goplantserver/middlewares"
)

func main() {
	router := http.NewServeMux()

	// Public routes
	router.Handle("/", middlewares.UseGlobalMiddlewares(handlers.HelloHandler))
	router.Handle("/about", middlewares.UseGlobalMiddlewares(handlers.AboutHandler))

	// Protected routes
	router.Handle("/plants", middlewares.ProtectRoute(handlers.PlantsHandler))

	// Start server
	logger.Logger.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Logger.Fatal("Error starting server:", err)
	}
}
