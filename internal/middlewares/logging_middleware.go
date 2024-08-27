package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware that logs the start and end time of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request method and path before processing the request
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the time taken to complete the request after processing
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
