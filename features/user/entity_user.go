package user

import "mime/multipart"

type RegisterCore struct {
	ID       uint
	UserName string
	Password string
	Status   string
	Role     int
	FullName string
	Email    string
	Notelpn  string
}

type LoginCore struct {
	ID       uint
	UserName string
	Password string
	Status   string
}

type InsertCustomer struct {
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

type UserDataInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(username, password string) (data *LoginCore, err error)
	VerifiedEmail(id uint, input RegisterCore) error
	InsertCustomer(input InsertCustomer, file multipart.File, nameFile string) error
}

type UserServiceInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(username, password string) (data *LoginCore, token string, err error)
	VerifiedEmail(id uint, input RegisterCore) error
	InsertCustomer(input InsertCustomer, file multipart.File, nameFile string) error
}
