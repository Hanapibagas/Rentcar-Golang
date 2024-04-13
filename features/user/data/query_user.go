package data

import (
	"StartUp-Go/app/database"
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/user"
	"StartUp-Go/utils/encrypts"
	"StartUp-Go/utils/uploads"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userQuery struct {
	db            *gorm.DB
	hashService   encrypts.HashInterface
	uploadService uploads.CloudinaryInterface
}

func NewUser(db *gorm.DB, hash encrypts.HashInterface, cloud uploads.CloudinaryInterface) user.UserDataInterface {
	return &userQuery{
		db:            db,
		hashService:   hash,
		uploadService: cloud,
	}
}

func (repo *userQuery) InsertCustomer(input user.InsertCustomer, file multipart.File, nameFile string) error {
	var folderName string = "img/pelanggan"
	uuid := uuid.New().String()
	hashedPassword, err := repo.hashService.HashPassword(input.Password)
	if err != nil {
		return errors.New("error hashing password")
	}

	if file != nil {
		imgUrl, errUpload := repo.uploadService.Upload(file, nameFile, folderName)
		if errUpload != nil {
			return errors.New("error upload img")
		}

		input.FotoKtp = imgUrl.SecureURL
	}

	inputRegisterGorm := database.User{
		Uuid:     uuid,
		UserName: input.UserName,
		Password: hashedPassword,
		Status:   "F",
		Role:     input.Role,
	}

	repo.db.Create(&inputRegisterGorm)

	inputBiodataGorm := database.Biodata{
		UuidUser:      inputRegisterGorm.Uuid,
		Uuid:          uuid,
		FullName:      input.FullName,
		TempatLahir:   input.TempatLahir,
		Alamat:        input.Alamat,
		Email:         input.Email,
		Notelp:        input.Notelp,
		NotelpKerabat: input.NotelpKerabat,
		Ktp:           input.Ktp,
		Pekerjaan:     input.Pekerjaan,
		FotoKtp:       input.FotoKtp,
	}

	repo.db.Create(&inputBiodataGorm)

	return nil
}

func (repo *userQuery) VerifiedEmail(id uint, input user.RegisterCore) error {
	userEmailVerified := database.User{
		Status: database.UserRole(input.Status),
	}

	tx := repo.db.Model(&database.User{}).Where("id = ?", id).Updates(&userEmailVerified)

	if tx.Error != nil {
		return errors.New("database error: " + tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return errors.New("no rows affected, user not found")
	}

	return nil
}

func (repo *userQuery) Login(username string, password string) (data *user.LoginCore, err error) {
	var LoginData database.User
	tx := repo.db.Where("user_name = ?", username).First(&LoginData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := LoginData.ModelToLogin()

	return &result, nil
}

func (repo *userQuery) Register(input user.RegisterCore) (data *user.RegisterCore, token string, err error) {
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
