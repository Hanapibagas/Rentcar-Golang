package service

import (
	"StartUp-Go/app/middlewares"
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

func (service *userService) VerifiedEmail(id int, input user.UserCore) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.userData.VerifiedEmail(id, input)
	return err
}

func (service *userService) Login(email string, password string) (data *user.UserCore, token string, err error) {
	if email == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}

	data, err = service.userData.Login(email, password)
	if err != nil {
		return nil, "", errors.New("Email atau password salah")
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}
	return data, token, err
}

func (service *userService) Register(input user.UserCore) (data *user.UserCore, token string, err error) {
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
