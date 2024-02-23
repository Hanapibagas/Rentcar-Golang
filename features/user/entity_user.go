package user

import "mime/multipart"

type UserCore struct {
	ID          uint   `json:"id"`
	UserName    string `json:"user_name" validate:"required"`
	ShopName    string `json:"shop_name"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number"`
	Domicile    string `json:"domicile" validate:"required"`
	Address     string `json:"address"`
	Image       string `json:"image"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Subdistrict string `json:"subdistrict"`
	Tagline     string `json:"tagline"`
	ShopImage   string `json:"shop_image"`
}

type UserDataInterface interface {
	SelectUser(id int) (*UserCore, error)
	SelectShop(id int) (*UserCore, error)
	UpdateUser(id int, input UserCore, file multipart.File, nameFile string) error
	UpdateShop(id int, input UserCore, file multipart.File, nameFile string) error
	Delete(id int) error
}

type UserServiceInterface interface {
	SelectUser(id int) (*UserCore, error)
	SelectShop(id int) (*UserCore, error)
	UpdateUser(id int, input UserCore, file multipart.File, nameFile string) error
	UpdateShop(id int, input UserCore, file multipart.File, nameFile string) error
	Delete(id int) error
}
