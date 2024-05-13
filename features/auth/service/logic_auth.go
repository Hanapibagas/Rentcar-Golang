package service

import (
	"StartUp-Go/features/auth"
	"StartUp-Go/utils/encrypts"
	"errors"

	"github.com/go-playground/validator/v10"
)

type authService struct {
	authData    auth.AuthDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

func NewAuth(repo auth.AuthDataInterface, hash encrypts.HashInterface) auth.AuthServiceInterface {
	return &authService{
		authData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

func (a *authService) Login(username string, password string) (data *auth.LoginCore, token string, err error) {
	panic("unimplemented")
}

func (service *authService) Register(input auth.RegisterCore) (data *auth.RegisterCore, token string, err error) {
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

	data, generatedToken, err := service.authData.Register(input)
	return data, generatedToken, err
}

// VerifiedEmail implements auth.AuthServiceInterface.
func (a *authService) VerifiedEmail(id uint, input auth.RegisterCore) error {
	panic("unimplemented")
}
