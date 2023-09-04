package router

import (
	"kelompok1/immersive-dash/features/user/data"
	"kelompok1/immersive-dash/features/user/handler"
	"kelompok1/immersive-dash/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := data.New(db)
	userService := service.New(userRepo)
	userHandlerAPI := handler.New(userService)

	e.POST("/login", userHandlerAPI.Login)
}
