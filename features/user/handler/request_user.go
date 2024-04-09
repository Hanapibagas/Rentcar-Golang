package handler

import "StartUp-Go/features/user"

type UserRequestRegister struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Status   string
	Role     string
	Email    string `json:"email"`
	Notelp   string `json:"no_telp"`
}

type UserRequestLogin struct {
	Email    string
	Password string
}

func RequestUserRegisterToCore(input UserRequestRegister) user.RegisterCore {
	return user.RegisterCore{
		UserName: input.UserName,
		Password: input.Password,
		Email:    input.Email,
		Notelpn:  input.Notelp,
	}
}
