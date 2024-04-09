package routes

import (
	_authData "StartUp-Go/features/user/data"
	_authHandler "StartUp-Go/features/user/handler"
	_authService "StartUp-Go/features/user/service"
	"StartUp-Go/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()

	authData := _authData.NewUser(db)
	autService := _authService.NewUser(authData, hashService)
	authHandler := _authHandler.NewUser(autService)

	// login
	e.POST("/register", authHandler.RegisterUser)
	// 	e.POST("/login", authHandler.LoginUser)
	// 	e.POST("/verified", authHandler.VerifiedEmail, middlewares.JWTMiddleware())
	// 	e.PUT("/update-password", authHandler.UpdatePassword, middlewares.JWTMiddleware())
}
