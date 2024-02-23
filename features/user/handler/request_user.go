package handler

import "StartUp-Go/features/user"

type UserRequestRegister struct {
	Name       string
	Occupation string
	Email      string
	Password   string
	Role       string
}

func RequestUserRegisterToCore(input UserRequestRegister) user.UserCore {
	role := "user"
	if input.Role != "" {
		role = input.Role
	}
	return user.UserCore{
		Name:       input.Name,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   input.Password,
		Role:       role,
	}
}
