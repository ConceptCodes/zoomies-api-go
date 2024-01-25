package helpers

import (
	"encoding/json"
	"net/http"
	"zoomies-api-go/pkg/models"
)

func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := models.Response{
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func SendErrorResponse(w http.ResponseWriter, message string, errorCode string) {

	response := models.Response{
		Message:   message,
		ErrorCode: errorCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	switch errorCode {
	case "ZM-404":
		w.WriteHeader(http.StatusNotFound)
	case "ZM-401":
		w.WriteHeader(http.StatusUnauthorized)
	case "ZM-403":
		w.WriteHeader(http.StatusForbidden)
	case "ZM-500":
		w.WriteHeader(http.StatusInternalServerError)
	}
}
