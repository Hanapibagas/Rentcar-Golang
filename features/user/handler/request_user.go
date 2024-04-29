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
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserRequestVerified struct {
	Status string `json:"status"`
}

type InsertCostumer struct {
	UserName      string `json:"user_name" form:"user_name"`
	Password      string `json:"password" form:"password"`
	Status        string `json:"status" form:"status"`
	Role          int    `json:"role" form:"role"`
	FullName      string `json:"full_name" form:"full_name"`
	TempatLahir   string `json:"tempat_lahir" form:"tempat_lahir"`
	Alamat        string `json:"alamat" form:"alamat"`
	Email         string `json:"email" form:"email"`
	Notelp        string `json:"notelp" form:"notelp"`
	NotelpKerabat string `json:"notelp_kerabat" form:"notelp_kerabat"`
	Ktp           string `json:"ktp" form:"ktp"`
	Pekerjaan     string `json:"pekerjaan" form:"pekerjaan"`
	FotoKtp       string `json:"foto_ktp" form:"foto_ktp"`
}

type InsertUser struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Status   string `json:"status" form:"status"`
	Role     string `json:"role" form:"role"`
	FullName string `json:"full_name" form:"full_name"`
	Alamat   string `json:"alamat" form:"alamat"`
	Email    string `json:"email" form:"email"`
	Notelp   string `json:"notelp" form:"notelp"`
	FotoKtp  string `json:"foto_ktp" form:"foto_ktp"`
}

func RequestToCoreUser(input InsertUser) user.InstertUser {
	return user.InstertUser{
		UserName: input.UserName,
		Password: input.Password,
		Status:   input.Status,
		Role:     input.Role,
		FullName: input.FullName,
		Alamat:   input.Alamat,
		Email:    input.Email,
		Notelp:   input.Notelp,
		FotoKtp:  input.FotoKtp,
	}
}

type UpdatetCostumer struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	// Status        string `json:"status" form:"status"`
	// Role          int    `json:"role" form:"role"`
	FullName      string `json:"full_name" form:"full_name"`
	TempatLahir   string `json:"tempat_lahir" form:"tempat_lahir"`
	Alamat        string `json:"alamat" form:"alamat"`
	Email         string `json:"email" form:"email"`
	Notelp        string `json:"notelp" form:"notelp"`
	NotelpKerabat string `json:"notelp_kerabat" form:"notelp_kerabat"`
	Ktp           string `json:"ktp" form:"ktp"`
	Pekerjaan     string `json:"pekerjaan" form:"pekerjaan"`
	FotoKtp       string `json:"foto_ktp" form:"foto_ktp"`
}

func RequestToUpdateCostumer(input UpdatetCostumer) user.UpdateCustomer {
	return user.UpdateCustomer{
		UserName: input.UserName,
		Password: input.Password,
		// Status:        input.Status,
		// Role:          input.Role,
		FullName:      input.FullName,
		TempatLahir:   input.TempatLahir,
		Alamat:        input.Alamat,
		Email:         input.Email,
		Notelp:        input.Notelp,
		NotelpKerabat: input.NotelpKerabat,
		Ktp:           input.Ktp,
		Pekerjaan:     input.Pekerjaan,
		FotoKtp:       input.FotoKtp,
	}
}
func RequestToInsertCostumer(input InsertCostumer) user.InsertCustomer {
	return user.InsertCustomer{
		UserName:      input.UserName,
		Password:      input.Password,
		Status:        input.Status,
		Role:          input.Role,
		FullName:      input.FullName,
		TempatLahir:   input.TempatLahir,
		Alamat:        input.Alamat,
		Email:         input.Email,
		Notelp:        input.Notelp,
		NotelpKerabat: input.NotelpKerabat,
		Ktp:           input.Ktp,
		Pekerjaan:     input.Pekerjaan,
		FotoKtp:       input.FotoKtp,
	}
}

func RequestToUpdateVerified(input UserRequestVerified) user.RegisterCore {
	return user.RegisterCore{
		Status: input.Status,
	}
}

func RequestUserRegisterToCore(input UserRequestRegister) user.RegisterCore {
	return user.RegisterCore{
		UserName: input.UserName,
		Password: input.Password,
		Email:    input.Email,
		Notelpn:  input.Notelp,
	}
}
