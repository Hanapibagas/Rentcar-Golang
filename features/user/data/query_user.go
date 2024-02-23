package data

import (
	"BE-REPO-20/features/user"
	"BE-REPO-20/utils/uploads"
	"errors"
	"mime/multipart"
	"strings"

	"gorm.io/gorm"
)

type userQuery struct {
	db            *gorm.DB
	uploadService uploads.CloudinaryInterface
}

func NewUser(db *gorm.DB, cloud uploads.CloudinaryInterface) user.UserDataInterface {
	return &userQuery{
		db:            db,
		uploadService: cloud,
	}
}

func (repo *userQuery) GetUserImageById(userId int) (string, error) {
	var user User
	if err := repo.db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		return "", err
	}

	return user.Image, nil
}

func (repo *userQuery) GetUserShopImageById(userId int) (string, error) {
	var user User
	if err := repo.db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		return "", err
	}

	return user.ShopImage, nil
}

// SelectShop implements user.UserDataInterface.
func (repo *userQuery) SelectShop(id int) (*user.UserCore, error) {
	var userGorm User
	tx := repo.db.Where("id = ?", id).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := userGorm.ModelToCoreShop()
	return &result, nil
}

// SelectUser implements user.UserDataInterface.
func (repo *userQuery) SelectUser(id int) (*user.UserCore, error) {
	var userGorm User
	tx := repo.db.Where("id = ?", id).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := userGorm.ModelToCore()
	return &result, nil
}

// UpdateShop implements user.UserDataInterface.
func (repo *userQuery) UpdateShop(id int, input user.UserCore, file multipart.File, nameFile string) error {
	// foldering name at cloudinary
	var folderName string = "img/users"

	// get user shop image from db
	imgGorm, _ := repo.GetUserShopImageById(id)

	// if shop image is not nil destroy the old image from cloudinary
	if imgGorm != "" {
		splitImgSlash := strings.Split(imgGorm, "/")
		publicIdSliceSlash := splitImgSlash[7:10]
		publicIdGormSlash := strings.Join(publicIdSliceSlash, "/")

		splitPublicId := strings.Split(publicIdGormSlash, ".")
		publicIdSliced := splitPublicId[0:(len(splitPublicId) - 1)]
		publicId := strings.Join(publicIdSliced, ".")

		repo.uploadService.Destroy(publicId)
	}

	// upload file to cloudinary and input to db
	if file != nil {
		imgUrl, errUpload := repo.uploadService.Upload(file, nameFile, folderName)
		if errUpload != nil {
			return errors.New("error upload img")
		}

		input.ShopImage = imgUrl.SecureURL
	}
	userGorm := CoreToModelShop(input)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}

// UpdateUser implements user.UserDataInterface.
func (repo *userQuery) UpdateUser(id int, input user.UserCore, file multipart.File, nameFile string) error {
	// foldering name at cloudinary
	var folderName string = "img/users"

	// get user image from db
	imgGorm, _ := repo.GetUserImageById(id)

	// if user image is not nil destroy the old image from cloudinary
	if imgGorm != "" {
		splitImgSlash := strings.Split(imgGorm, "/")
		publicIdSliceSlash := splitImgSlash[7:10]
		publicIdGormSlash := strings.Join(publicIdSliceSlash, "/")

		splitPublicId := strings.Split(publicIdGormSlash, ".")
		publicIdSliced := splitPublicId[0:(len(splitPublicId) - 1)]
		publicId := strings.Join(publicIdSliced, ".")

		repo.uploadService.Destroy(publicId)
	}

	// upload file to cloudinary and input to db
	if file != nil {
		imgUrl, errUpload := repo.uploadService.Upload(file, nameFile, folderName)
		if errUpload != nil {
			return errors.New("error upload img")
		}
		input.Image = imgUrl.SecureURL
	}

	userGorm := CoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(id int) error {
	tx := repo.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}
