package middlewares

import (
	"net/http"
	"os"

	"github.com/jasonflorentino/goplantserver/logger"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Printf("Incoming request: %s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Cool-Header", "My Cool Header")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		providedkey := r.FormValue("api_key")
		expectedkey := os.Getenv("PLANTSERVER_API_KEY")

		if expectedkey == "" {
			logger.Logger.Printf("Error: No API Key set!")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		if providedkey != expectedkey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func UseGlobalMiddlewares(handler http.HandlerFunc) http.Handler {
	return LoggerMiddleware(HeaderMiddleware(handler))
}

func ProtectRoute(handler http.HandlerFunc) http.Handler {
	return LoggerMiddleware(HeaderMiddleware(AuthMiddleware(handler)))
}
