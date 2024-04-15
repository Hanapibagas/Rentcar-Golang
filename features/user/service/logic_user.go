package service

import (
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/user"
	"StartUp-Go/utils/encrypts"
	"errors"
	"mime/multipart"

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

func (service *userService) DeleteByUuid(uuid string) error {
	err := service.userData.DeleteByUuid(uuid)
	return err
}

func (service *userService) GetByUuid(uuid string) (*user.GetByIdCustomer, error) {
	result, err := service.userData.GetByUuid(uuid)
	return result, err
}

func (service *userService) GetAllCostumer() ([]user.GetAllCustomer, error) {
	resul, err := service.userData.GetAllCostumer()
	return resul, err
}

func (service *userService) UpdateCustomer(uuid string, input user.UpdateCustomer, file multipart.File, nameFile string) error {
	if uuid == "" {
		return errors.New("invalid uuid")
	}

	err := service.userData.UpdateCustomer(uuid, input, file, nameFile)
	return err
}

func (service *userService) InsertCustomer(input user.InsertCustomer, file multipart.File, nameFile string) error {
	err := service.userData.InsertCustomer(input, file, nameFile)
	return err
}

func (service *userService) VerifiedEmail(id uint, input user.RegisterCore) error {
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

func (service *userService) Login(username string, password string) (data *user.LoginCore, token string, err error) {
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
