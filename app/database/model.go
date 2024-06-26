package database

import (
	"StartUp-Go/features/auth"
	"time"
)

type UserRole string
type GenderBiodata string

const (
	True  UserRole = "T"
	False UserRole = "F"
)

const (
	LakiLaki  GenderBiodata = "L"
	Perempuan GenderBiodata = "P"
)

type User struct {
	// gorm.Model
	ID       uint     `gorm:"primaryKey"`
	Uuid     string   `gorm:"type:char(38);default:null"`
	UserName string   `gorm:"size:100;default:null"`
	Password string   `gorm:"size:255;default:null"`
	Status   UserRole `gorm:"type:ENUM('T','F');default:null"`
	Role     int      `gorm:"default:null"`
}

type Biodata struct {
	// gorm.Model
	ID            uint          `gorm:"primaryKey"`
	UuidUser      string        `gorm:"size:38;default:null"`
	Uuid          string        `gorm:"type:char(38);default:null"`
	FullName      string        `gorm:"size:128;default:null"`
	Gender        GenderBiodata `gorm:"type:ENUM('L','P');default:null"`
	TempatLahir   string        `gorm:"size:128;default:null"`
	TanggalLahir  time.Time     `gorm:"default:null"`
	Alamat        string        `gorm:"type:LONGTEXT;default:null"`
	Email         string        `gorm:"size:100;default:null"`
	Notelp        string        `gorm:"size:20;default:null"`
	NotelpKerabat string        `gorm:"size:20;default:null"`
	Profile       string        `gorm:"size:128;default:null"`
	Ktp           string        `gorm:"size:128;default:null"`
	Pekerjaan     string        `gorm:"size:128;default:null"`
	FotoKtp       string        `gorm:"size:128;default:null"`
}

// func CoreToListCatomer

// func CoreUpdate(input user.InsertCustomer) Biodata {
// 	return Biodata{
// 		FullName:      input.FullName,
// 		TempatLahir:   input.TempatLahir,
// 		Alamat:        input.Alamat,
// 		Email:         input.Email,
// 		Notelp:        input.Notelp,
// 		NotelpKerabat: input.NotelpKerabat,
// 		Ktp:           input.Ktp,
// 		Pekerjaan:     input.Pekerjaan,
// 		FotoKtp:       input.FotoKtp,
// 	}
// }

// func (u User) ModelToLogin() auth.LoginCore {
// 	return auth.LoginCore{
// 		ID:       u.ID,
// 		Uuid:     u.Uuid,
// 		UserName: u.UserName,
// 		Status:   string(u.Status),
// 		Password: u.Password,
// 	}
// }

func (u User) ModelToCore() auth.RegisterCore {
	return auth.RegisterCore{
		UserName: u.UserName,
		Password: u.Password,
	}
}
