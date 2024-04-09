package data

import (
	"StartUp-Go/app/database"
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func (repo *userQuery) Register(input user.RegisterCore) (data *user.RegisterCore, token string, err error) {
	uuid := uuid.New().String()
	inputRegisterGorm := database.User{
		Uuid:     uuid,
		UserName: input.UserName,
		Password: input.Password,
		Status:   "F",
		Role:     1,
	}

	tx := repo.db.Create(&inputRegisterGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	inputBiodataGorm := database.Biodata{
		UuidUser: inputRegisterGorm.Uuid,
		Email:    input.Email,
		Notelp:   input.Notelpn,
	}
	if err := repo.db.Create(&inputBiodataGorm).Error; err != nil {
		return nil, "", err
	}

	var authGorm database.User
	result := authGorm.ModelToCore()

	generatedToken, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return nil, "", err
	}

	return &result, generatedToken, nil
}
