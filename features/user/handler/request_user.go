package handler

import "StartUp-Go/features/user"

type UserRequestRegister struct {
	Name              string
	Occupation        string
	Email             string
	Password          string
	EmailVerification string
	Role              string
}


type UserRequestLogin struct {
	Email    string
	Password string
}

type UserRequestVerified struct {
	EmailVerification string `json:"email_verified"`
}

func RequestToUpdateVerified(input UserRequestVerified) user.EmailVerification {
	return user.EmailVerification{
		EmailVerification: input.EmailVerification,
	}
}

func RequestUserRegisterToCore(input UserRequestRegister) user.UserCore {
	role := "user"
	if input.Role != "" {
		role = input.Role
	}
	Verification := "not yet verified"
	if input.EmailVerification != "" {
		Verification = input.EmailVerification
	}
	return user.UserCore{
		ID:                0,
		Name:              input.Name,
		Occupation:        input.Occupation,
		Email:             input.Email,
		EmailVerification: Verification,
		Password:          input.Password,
		Role:              role,
	}
}
