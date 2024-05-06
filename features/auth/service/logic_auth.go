package service

import (
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/auth"
	"StartUp-Go/utils/encrypts"
	"errors"

	"github.com/go-playground/validator/v10"
)

type authrService struct {
	userData    auth.AuthDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

func NewUser(repo auth.AuthDataInterface, hash encrypts.HashInterface) auth.AuthServiceInterfave {
	return &authrService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

func (service *authrService) Login(username string, password string) (data *auth.LoginCore, token string, err error) {
	if username == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}

	data, err = service.userData.Login(username, password)
	if err != nil {
		return nil, "", errors.New("email atau password salah")
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}

	return data, token, err
}

func (service *authrService) Register(input auth.RegisterCore) (data *auth.RegisterCore, token string, err error) {
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

func (service *authrService) VerifiedEmail(id uint, input auth.RegisterCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	// Periksa apakah EmailVerification tidak kosong
	if input.Status == "" {
		return errors.New("Status Verification is empty.")
	}

	err := service.userData.VerifiedEmail(id, input)
	return err
}
