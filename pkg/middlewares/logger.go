package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Started %s %s", r.Method, r.URL.Path)

		if r.Method == "POST" || r.Method == "PUT" {
			r.ParseForm()
			log.Printf("Request Body: %s", r.Form)
		}

		next.ServeHTTP(w, r)

		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("Response Body: %s", w)

		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
