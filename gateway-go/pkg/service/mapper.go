package service

import (
	"ledungcobra/gateway-go/pkg/controllers/users/request"
	"ledungcobra/gateway-go/pkg/models"
)

func mapRegisterRequestToUser(request request.RegisterRequest) models.User {
	return models.User{
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Email:      request.Email,
		BirthDay:   request.BirthDay,
		BirthYear:  request.BirthYear,
		BirthMonth: request.BirthMonth,
		Gender:     request.Gender,
	}
}
