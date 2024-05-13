package handler

import "StartUp-Go/features/auth"

type UserRequestRegister struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Status   string
	Role     string
	Email    string `json:"email"`
	Notelp   string `json:"no_telp"`
}

func RequestUserRegisterToCore(input UserRequestRegister) auth.RegisterCore {
	return auth.RegisterCore{
		UserName: input.UserName,
		Password: input.Password,
		Email:    input.Email,
		Notelpn:  input.Notelp,
	}
}
