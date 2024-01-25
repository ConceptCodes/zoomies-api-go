package middlewares

import (
	"log"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%s %s", r.Method, r.URL.Path)

		if r.Method == "POST" || r.Method == "PUT" {
			r.ParseForm()
			log.Printf("Request Body: %s", r.Form)
		}

		next.ServeHTTP(w, r)

	})
}

func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

		log.Printf("Response Body: %s", r.Body)

	})
}
