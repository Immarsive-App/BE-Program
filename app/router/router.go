package router

import (
	"kelompok1/immersive-dash/features/team/data"
	"kelompok1/immersive-dash/features/team/handler"
	"kelompok1/immersive-dash/features/team/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := data.New(db)
	userService := service.New(userRepo)
	userHandlerAPI := handler.New(userService)

	e.POST("/login", userHandlerAPI.Login)
}
