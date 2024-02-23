package database

import (
	"BE-REPO-20/app/configs"
	"BE-REPO-20/features/admin/data"
	_cartData "BE-REPO-20/features/cart/data"
	_productData "BE-REPO-20/features/product/data"

	// "BE-REPO-20/features/auth/data"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *configs.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return DB
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&data.User{})
	db.AutoMigrate(&data.Order{})
	db.AutoMigrate(&data.ItemOrder{})
	db.AutoMigrate(&_productData.Product{})
	db.AutoMigrate(&_productData.ProductImage{})
	db.AutoMigrate(&_cartData.Cart{})
}
