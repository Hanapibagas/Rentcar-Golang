package service

import (
	"BE-REPO-20/features/user"
	"errors"
	"mime/multipart"
)

type userService struct {
	userData user.UserDataInterface
}

func NewUser(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

// SelectShop implements user.UserServiceInterface.
func (service *userService) SelectShop(id int) (*user.UserCore, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	data, err := service.userData.SelectShop(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SelectUser implements user.UserServiceInterface.
func (service *userService) SelectUser(id int) (*user.UserCore, error) {
	data, err := service.userData.SelectUser(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateShop implements user.UserServiceInterface.
func (service *userService) UpdateShop(id int, input user.UserCore, file multipart.File, nameFile string) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.userData.UpdateShop(id, input, file, nameFile)
	return err
}

// UpdateUser implements user.UserServiceInterface.
func (service *userService) UpdateUser(id int, input user.UserCore, file multipart.File, nameFile string) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.userData.UpdateUser(id, input, file, nameFile)
	return err
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	err := service.userData.Delete(id)
	return err
}
