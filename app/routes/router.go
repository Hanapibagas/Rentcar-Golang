package routes

import (
	_authData "StartUp-Go/features/auth/data"
	_authHandler "StartUp-Go/features/auth/handler"
	_authService "StartUp-Go/features/auth/service"
	"StartUp-Go/utils/encrypts"
	"StartUp-Go/utils/uploads"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()
	uploadService := uploads.NewCloudService()

	authData := _authData.NewAuth(db, hashService, uploadService)
	autService := _authService.NewAuth(authData, hashService)
	authHandler := _authHandler.NewAuth(autService)

	// login
	e.POST("/register", authHandler.RegisterUser)
	// e.POST("/login", authHandler.LoginUser)
	// e.POST("/verified", authHandler.VerifiedEmail, middlewares.JWTMiddleware())

	// e.POST("/insert-costumer", authHandler.InsertCostumer, middlewares.JWTMiddleware())
	// e.POST("/insert-user", authHandler.InsertUser, middlewares.JWTMiddleware())
	// e.PUT("/update-costumer/:uuid", authHandler.UpdateCostumer, middlewares.JWTMiddleware())
	// e.GET("/costumer", authHandler.GetAllCostumer, middlewares.JWTMiddleware())
	// e.GET("/costumer/:uuid", authHandler.GetByUuidCostumer, middlewares.JWTMiddleware())
	// e.DELETE("/costumer/:uuid", authHandler.DeleteByUuidCostumer, middlewares.JWTMiddleware())
	// 	e.PUT("/update-password", authHandler.UpdatePassword, middlewares.JWTMiddleware())
}
