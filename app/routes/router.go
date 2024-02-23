package routes

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/utils/encrypts"
	"BE-REPO-20/utils/midtrans"
	"BE-REPO-20/utils/uploads"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	// login
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)
	e.PUT("/update-password", authHandler.UpdatePassword, middlewares.JWTMiddleware())
}
