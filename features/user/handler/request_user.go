package handler

import "BE-REPO-20/features/user"

type UserRequest struct {
	UserName    string `json:"user_name" form:"username"`
	Domicile    string `json:"domicile" form:"domicile"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Image       string `json:"image_url" form:"image"`
}

type UserShopRequest struct {
	ShopName    string `json:"shop_name" form:"shopname"`
	Tagline     string `json:"tagline" form:"tagline"`
	Province    string `json:"province" form:"province"`
	City        string `json:"city" form:"city"`
	Subdistrict string `json:"subdistrict" form:"subdistrict"`
	Address     string `json:"address" form:"address"`
	ShopImage   string `json:"shop_image" form:"shopimage"`
	Category    string `json:"category" form:"category"`
}

func RequestToCore(input UserRequest) user.UserCore {
	return user.UserCore{
		UserName:    input.UserName,
		Domicile:    input.Domicile,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Image:       input.Image,
	}
}

func RequestToCoreShop(input UserShopRequest) user.UserCore {
	return user.UserCore{
		ShopName:    input.ShopName,
		Tagline:     input.Tagline,
		Province:    input.Province,
		City:        input.City,
		Subdistrict: input.Subdistrict,
		Address:     input.Address,
		ShopImage:   input.ShopImage,
	}
}
