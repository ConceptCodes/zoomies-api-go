package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"zoomies-api-go/pkg/constants"
	"zoomies-api-go/pkg/helpers"
	"zoomies-api-go/pkg/models"
	"zoomies-api-go/pkg/repository"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.LoginDto
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.BadRequest)
	}

	storedUser, err := h.userRepo.FindByEmail(user.Email)

	if err != nil {
		message := fmt.Sprintf(constants.EntityNotFound, "User", user.Email)
		helpers.SendErrorResponse(w, message, constants.NotFound)
	}

	if helpers.CheckPasswordHash(user.Password, storedUser.Password) {
		helpers.SendErrorResponse(w, constants.LoginFailure, constants.Unauthorized)
	}

	token, err := helpers.GenerateToken(storedUser.ID, storedUser.Email)

	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.InternalServerError)
	}

	w.Header().Set("Authorization", "Bearer "+token)
	helpers.SendSuccessResponse(w, constants.LoginSuccessful, helpers.GetSimpleUser(*storedUser))
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.BadRequest)
	}

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.InternalServerError)
	}

	user.Password = hashedPassword

	err = h.userRepo.Save(&user)

	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.InternalServerError)
	}

	helpers.SendSuccessResponse(w, constants.RegistrationSuccessful, helpers.GetSimpleUser(user))
}
