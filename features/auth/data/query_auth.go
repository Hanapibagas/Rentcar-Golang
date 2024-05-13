package data

import (
	"StartUp-Go/app/database"
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/auth"
	"StartUp-Go/utils/encrypts"
	"StartUp-Go/utils/uploads"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authQuery struct {
	db            *gorm.DB
	hashService   encrypts.HashInterface
	uploadService uploads.CloudinaryInterface
}

func NewAuth(db *gorm.DB, hash encrypts.HashInterface, cloud uploads.CloudinaryInterface) auth.AuthDataInterface {
	return &authQuery{
		db:            db,
		hashService:   hash,
		uploadService: cloud,
	}
}

func (a *authQuery) Login(username string, password string) (data *auth.LoginCore, err error) {
	panic("unimplemented")
}

// Register implements auth.AuthDataInterface.
func (repo *authQuery) Register(input auth.RegisterCore) (data *auth.RegisterCore, token string, err error) {
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
		Uuid:     uuid,
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

// VerifiedEmail implements auth.AuthDataInterface.
func (a *authQuery) VerifiedEmail(id uint, input auth.RegisterCore) error {
	panic("unimplemented")
}
