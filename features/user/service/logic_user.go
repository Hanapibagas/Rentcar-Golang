package service

import (
	"StartUp-Go/features/user"
	"StartUp-Go/utils/encrypts"
	"errors"

	"github.com/go-playground/validator"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

func NewUser(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

func (service *userService) Register(input user.RegisterCore) (data *user.RegisterCore, token string, err error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return nil, "", errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return nil, "", errors.New("rror hashing password")
		}
		input.Password = hashedPass
	}

	data, generatedToken, err := service.userData.Register(input)
	return data, generatedToken, err
}
