package data

import (
	"BE-REPO-20/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"default:null" db:"user_name"`
	ShopName    string `gorm:"default:null"`
	Email       string `gorm:"default:null;unique"`
	PhoneNumber string `gorm:"default:null"`
	Domicile    string `gorm:"default:null"`
	Password    string `gorm:"default:null"`
	Address     string `gorm:"default:null"`
	Role        string `gorm:"default:user"`
	Image       string `gorm:"default:null"`
	Province    string `gorm:"default:null"`
	City        string `gorm:"default:null"`
	Subdistrict string `gorm:"default:null"`
	TagLine     string `gorm:"default:null"`
	ShopImage   string `gorm:"default:null"`
	Category    string `gorm:"default:null"`
}

func (u User) ModelToCore() user.UserCore {
	return user.UserCore{
		ID:          u.ID,
		UserName:    u.UserName,
		Email:       u.Email,
		Domicile:    u.Domicile,
		PhoneNumber: u.PhoneNumber,
		Image:       u.Image,
	}
}

func CoreToModel(input user.UserCore) User {
	return User{
		UserName:    input.UserName,
		Email:       input.Email,
		Domicile:    input.Domicile,
		PhoneNumber: input.PhoneNumber,
		Image:       input.Image,
	}
}

func CoreToModelShop(input user.UserCore) User {
	return User{
		ShopName:    input.ShopName,
		TagLine:     input.Tagline,
		Province:    input.Province,
		City:        input.City,
		Subdistrict: input.Subdistrict,
		Address:     input.Address,
		ShopImage:   input.ShopImage,
	}
}

func (u User) ModelToCoreShop() user.UserCore {
	return user.UserCore{
		ID:          u.ID,
		ShopName:    u.ShopName,
		Tagline:     u.TagLine,
		Province:    u.Province,
		City:        u.City,
		Subdistrict: u.Subdistrict,
		Address:     u.Address,
		ShopImage:   u.ShopImage,
	}
}
