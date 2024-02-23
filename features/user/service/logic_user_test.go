package service

import (
	"BE-REPO-20/features/user"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserShop(t *testing.T) {
	repo := new(mocks.UserData)
	srv := NewUser(repo)
	returnData := &user.UserCore{
		ID:          1,
		Province:    "Jakarta",
		City:        "Jakarta Timur",
		Subdistrict: "Duren Sawit",
		Tagline:     "Quality",
		ShopImage:   "www.image.com/photoshop",
	}
	t.Run("Succes Get User Shop", func(t *testing.T) {
		repo.On("SelectShop", 1, mock.Anything).Return(returnData, nil).Once()
		result, err := srv.SelectShop(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})

}

func TestGetUser(t *testing.T) {
	repo := new(mocks.UserData)
	srv := NewUser(repo)
	returnData := &user.UserCore{
		ID:          1,
		UserName:    "admin",
		ShopName:    "toko admin",
		Email:       "admin@mail.com",
		PhoneNumber: "0811111",
		Domicile:    "Jakarta",
		Address:     "Jl Soekarno",
		Image:       "www.image.com/photoshop",
	}
	t.Run("Succes Get User", func(t *testing.T) {
		repo.On("SelectUser", 1, mock.Anything).Return(returnData, nil).Once()
		result, err := srv.SelectUser(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)

		// userIdInvalid := 0
		// _, err = srv.SelectUser(1)
		// expectedErr := err
		// assert.EqualError(nil, err, expectedErr.Error())
	})
}

func TestUpdateUserShop(t *testing.T) {
	repo := new(mocks.UserData)
	srv := NewUser(repo)
	mockFile := new(multipart.File)
	input := user.UserCore{
		ID:          1,
		Province:    "Jakarta",
		City:        "Jakarta Timur",
		Subdistrict: "Duren Sawit",
		Tagline:     "Quality",
		ShopImage:   "www.image.com/photoshop",
	}
	t.Run("Succes Update User Shop", func(t *testing.T) {
		repo.On("UpdateShop", 1, input, *mockFile, "filename", mock.Anything).Return(nil).Once()
		err := srv.UpdateShop(1, input, *mockFile, "filename")
		assert.NoError(t, err)
		repo.AssertExpectations(t)

		userIdInvalid := 0
		err = srv.UpdateShop(userIdInvalid, input, *mockFile, "filename")
		expectedErr := errors.New("invalid id")
		assert.EqualError(t, err, expectedErr.Error())
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)
	srv := NewUser(repo)
	mockFile := new(multipart.File)
	input := user.UserCore{
		ID:          1,
		UserName:    "admin",
		Email:       "admin@mail.com",
		PhoneNumber: "0811111",
		Domicile:    "Jakarta",
		Address:     "Jl Soekarno",
		Image:       "www.image.com/photoshop",
	}
	t.Run("Succes Update User", func(t *testing.T) {
		repo.On("UpdateUser", 1, input, *mockFile, "filename", mock.Anything).Return(nil).Once()
		err := srv.UpdateUser(1, input, *mockFile, "filename")
		assert.NoError(t, err)
		repo.AssertExpectations(t)

		userIdInvalid := 0
		err = srv.UpdateUser(userIdInvalid, input, *mockFile, "filename")
		expectedErr := errors.New("invalid id")
		assert.EqualError(t, err, expectedErr.Error())
	})
}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)
	srv := NewUser(repo)
	userIdValid := 1
	userIdInvalid := 0
	t.Run("Succes Delete User", func(t *testing.T) {
		repo.On("Delete", 1, mock.Anything).Return(nil).Once()
		err := srv.Delete(userIdValid)
		assert.NoError(t, err)
		repo.AssertExpectations(t)

		err = srv.Delete(userIdInvalid)
		expectedErr := errors.New("invalid id")
		assert.EqualError(t, err, expectedErr.Error())
	})
}
