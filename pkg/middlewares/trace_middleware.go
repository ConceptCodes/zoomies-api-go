package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"zoomies-api-go/pkg/helpers"
	"zoomies-api-go/pkg/models"
)

func TraceRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestId := r.Header.Get("x-request-id")

		if requestId == "" {
			requestId = uuid.New().String()
		}

		w.Header().Add("x-request-id", requestId)

		authToken := r.Header.Get("Authorization")
		user := &models.SimpleUser{}

		if authToken != "" {
			authToken = strings.Replace(authToken, "Bearer ", "", 1)

			user, err := helpers.ValidateToken(authToken)

			if err != nil {
				log.Printf("Error: %s", err)
			}

			log.Printf("User Id: %d", user.ID)
		}

		log.Printf("Request Id: %s", requestId)

		model := &models.Request{
			Id:   requestId,
			User: *user,
		}

		ctx := context.WithValue(r.Context(), "ctx", model)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
