package handler

import "BE-REPO-20/features/user"

type UserResponse struct {
	UserName    string
	Domicile    string
	Email       string
	PhoneNumber string
	Image       string
}

type UserShopResponse struct {
	ShopName    string
	Tagline     string
	Province    string
	City        string
	Subdistrict string
	Address     string
	ShopImage   string
}

func CoreToResponse(data user.UserCore) UserResponse {
	return UserResponse{
		UserName:    data.UserName,
		Domicile:    data.Domicile,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Image:       data.Image,
	}
}
func CoreToResponseShop(data user.UserCore) UserShopResponse {
	return UserShopResponse{
		ShopName:    data.ShopName,
		Tagline:     data.Tagline,
		Province:    data.Province,
		City:        data.City,
		Subdistrict: data.Subdistrict,
		Address:     data.Address,
		ShopImage:   data.ShopImage,
	}
}
