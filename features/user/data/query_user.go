package data

import (
	"StartUp-Go/app/database"
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/user"
	"StartUp-Go/utils/encrypts"
	"StartUp-Go/utils/uploads"
	"errors"
	"mime/multipart"
	"strconv"

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

func (repo *userQuery) DeleteByUuid(uuid string) error {
	tx := repo.db.Begin()

	if err := tx.Where("uuid = ?", uuid).Delete(&database.User{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("uuid_user = ?", uuid).Delete(&database.Biodata{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (repo *userQuery) GetByUuid(uuid string) (*user.GetByIdCustomer, error) {
	var castomerUser database.User
	tx := repo.db.Where("uuid = ?", uuid).First(&castomerUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var biodata database.Biodata
	tx = repo.db.Where("uuid_user = ?", uuid).First(&biodata)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := &user.GetByIdCustomer{
		Uuid:          castomerUser.Uuid,
		UserName:      castomerUser.UserName,
		Status:        string(castomerUser.Status),
		Role:          strconv.Itoa(castomerUser.Role),
		FullName:      biodata.FullName,
		TempatLahir:   biodata.TempatLahir,
		Alamat:        biodata.Alamat,
		Email:         biodata.Email,
		Notelp:        biodata.Notelp,
		NotelpKerabat: biodata.NotelpKerabat,
		Ktp:           biodata.Ktp,
		Pekerjaan:     biodata.Pekerjaan,
		FotoKtp:       biodata.FotoKtp,
	}

	return result, nil
}

func (repo *userQuery) GetAllCostumer() ([]user.GetAllCustomer, error) {
	var castomerUser []database.User
	tx := repo.db.Where("role = ?", 5).Find(&castomerUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var castomerDataCore []user.GetAllCustomer
	for _, value := range castomerUser {
		var biodata []database.Biodata
		repo.db.Where("uuid_user = ?", value.Uuid).Find(&biodata)

		var biodataRespon user.GetAllCustomer
		if len(biodata) > 0 {
			biodataRespon = user.GetAllCustomer{
				Uuid:          value.Uuid,
				UserName:      value.UserName,
				Status:        string(value.Status),
				Role:          strconv.Itoa(value.Role),
				FullName:      biodata[0].FullName,
				TempatLahir:   biodata[0].TempatLahir,
				Alamat:        biodata[0].Alamat,
				Email:         biodata[0].Email,
				Notelp:        biodata[0].Notelp,
				NotelpKerabat: biodata[0].NotelpKerabat,
				Ktp:           biodata[0].Ktp,
				Pekerjaan:     biodata[0].Pekerjaan,
				FotoKtp:       biodata[0].FotoKtp,
			}
		}
		castomerDataCore = append(castomerDataCore, biodataRespon)
	}

	return castomerDataCore, nil
}

func (repo *userQuery) UpdateCustomer(uuid string, input user.UpdateCustomer, file multipart.File, nameFile string) error {
	var folderName string = "img/pelanggan"
	hashedPassword, err := repo.hashService.HashPassword(input.Password)
	if err != nil {
		return errors.New("error hashing password")
	}

	var userData database.User
	if err := repo.db.Where("uuid = ?", uuid).First(&userData).Error; err != nil {
		return errors.New("user not found")
	}

	inputRegisterGorm := database.User{
		UserName: input.UserName,
		Password: hashedPassword,
		// Status:   database.UserRole(input.Status),
		// Role:     input.Role,
	}
	repo.db.Model(&userData).Updates(&inputRegisterGorm)

	if file != nil {
		imgUrl, errUpload := repo.uploadService.Upload(file, nameFile, folderName)
		if errUpload != nil {
			return errors.New("error upload img")
		}
		input.FotoKtp = imgUrl.SecureURL
	}

	var biodataData database.Biodata
	if err := repo.db.Where("uuid_user = ?", userData.Uuid).First(&biodataData).Error; err != nil {
		return errors.New("biodata not found")
	}

	inputBiodataGorm := database.Biodata{
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
	repo.db.Model(&biodataData).Updates(&inputBiodataGorm)

	return nil
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
		Role:     5,
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

func (repo *userQuery) InsertUser(input user.InstertUser, file multipart.File, nameFile string) error {
	var folderName string = "img/user"
	uuid := uuid.New().String()

	if input.Password == "" {
		input.Password = "12345678"
	}
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

	roleInt, err := strconv.Atoi(input.Role)
	if err != nil {
		return errors.New("error converting role to int")
	}

	inputRegisterGorm := database.User{
		Uuid:     uuid,
		UserName: input.UserName,
		Password: hashedPassword,
		Status:   "F",
		Role:     roleInt,
	}

	repo.db.Create(&inputRegisterGorm)

	inputBiodataGorm := database.Biodata{
		UuidUser: inputRegisterGorm.Uuid,
		Uuid:     uuid,
		FullName: input.FullName,
		Alamat:   input.Alamat,
		Email:    input.Email,
		Notelp:   input.Notelp,
		FotoKtp:  input.FotoKtp,
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
