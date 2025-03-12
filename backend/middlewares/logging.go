package middlewares

import (
	"log"
	"net/http"
)

// const logentryCountMax int = 1000
// var logEntryCount int = 0

// func newLine() {
// 	logEntryCount++
// 	if logEntryCount > logentryCountMax {

// 	}

// }

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for '%s'", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
