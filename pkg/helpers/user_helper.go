package helpers

import (
	"strconv"
	"zoomies-api-go/pkg/models"
)

func GetSimpleUser(user models.User) models.SimpleUser {
	return models.SimpleUser{
		ID:        strconv.Itoa(int(user.ID)),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
