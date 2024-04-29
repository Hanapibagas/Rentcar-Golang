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
	Uuid     string
	Status   string
}

type UpdateCustomer struct {
	Uuid          string
	UserName      string
	Password      string
	Status        string
	Role          int
	FullName      string
	TempatLahir   string
	Alamat        string
	Email         string
	Notelp        string
	NotelpKerabat string
	Ktp           string
	Pekerjaan     string
	FotoKtp       string
}

type GetAllCustomer struct {
	Uuid          string
	UserName      string
	Password      string
	Status        string
	Role          string
	FullName      string
	TempatLahir   string
	Alamat        string
	Email         string
	Notelp        string
	NotelpKerabat string
	Ktp           string
	Pekerjaan     string
	FotoKtp       string
}

type GetByIdCustomer struct {
	Uuid          string
	UserName      string
	Password      string
	Status        string
	Role          string
	FullName      string
	TempatLahir   string
	Alamat        string
	Email         string
	Notelp        string
	NotelpKerabat string
	Ktp           string
	Pekerjaan     string
	FotoKtp       string
}

type InsertCustomer struct {
	UserName      string
	Password      string
	Status        string
	Role          int
	FullName      string
	TempatLahir   string
	Alamat        string
	Email         string
	Notelp        string
	NotelpKerabat string
	Ktp           string
	Pekerjaan     string
	FotoKtp       string
}
type InstertUser struct {
	Uuid     string
	FullName string
	UserName string
	Email    string
	Notelp   string
	Password string
	Alamat   string
	Role     string
	FotoKtp  string
	Status   string
}

type UserDataInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(username, password string) (data *LoginCore, err error)
	VerifiedEmail(id uint, input RegisterCore) error
	InsertCustomer(input InsertCustomer, file multipart.File, nameFile string) error
	InsertUser(input InstertUser, file multipart.File, nameFile string) error
	UpdateCustomer(uuid string, input UpdateCustomer, file multipart.File, nameFile string) error
	GetAllCostumer() ([]GetAllCustomer, error)
	GetByUuid(uuid string) (*GetByIdCustomer, error)
	DeleteByUuid(uuid string) error
}

type UserServiceInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(username, password string) (data *LoginCore, token string, err error)
	VerifiedEmail(id uint, input RegisterCore) error
	InsertCustomer(input InsertCustomer, file multipart.File, nameFile string) error
	InsertUser(input InstertUser, file multipart.File, nameFile string) error
	UpdateCustomer(uuid string, input UpdateCustomer, file multipart.File, nameFile string) error
	GetAllCostumer() ([]GetAllCustomer, error)
	GetByUuid(uuid string) (*GetByIdCustomer, error)
	DeleteByUuid(uuid string) error
}
