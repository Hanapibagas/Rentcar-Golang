package data

import (
	"StartUp-Go/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Occupation      string
	Email           string `gorm:"default:null;unique"`
	EmailVerification string
	AvatarFileName  string
	Password        string
	Role            string
}

func (u User) ModelToCore() user.UserCore {
	return user.UserCore{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Email:      u.Email,
		Password:   u.Password,
		Role:       u.Role,
	}
}
