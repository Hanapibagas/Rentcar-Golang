package data

import (
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/user"
	"errors"

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

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.UserCore, err error) {
	panic("unimplemented")
}

// Register implements user.UserDataInterface.
func (repo *userQuery) Register(input user.UserCore) (data *user.UserCore, token string, err error) {
	inputRegisterGorm := User{
		Name:       input.Name,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   input.Password,
		Role:       "user",
	}

	tx := repo.db.Create(&inputRegisterGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, "", errors.New("insert failed, row affected = 0")
	}

	var authGorm User
	tx = repo.db.Where("email = ?", input.Email).First(&authGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	result := authGorm.ModelToCore()

	generatedToken, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return nil, "", err
	}

	return &result, generatedToken, nil
}
