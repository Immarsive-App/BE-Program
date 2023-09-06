package router

import (
	"kelompok1/immersive-dash/app/middlewares"
	"kelompok1/immersive-dash/features/user/data"
	"kelompok1/immersive-dash/features/user/handler"
	"kelompok1/immersive-dash/features/user/service"

	_classData "kelompok1/immersive-dash/features/class/data"
	_classHandler "kelompok1/immersive-dash/features/class/handler"
	_classService "kelompok1/immersive-dash/features/class/service"

	_statusData "kelompok1/immersive-dash/features/status/data"
	_statusHandler "kelompok1/immersive-dash/features/status/handler"
	_statusService "kelompok1/immersive-dash/features/status/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := data.New(db)
	userService := service.New(userRepo)
	userHandlerAPI := handler.New(userService)

	classRepo := _classData.New(db)
	classService := _classService.New(classRepo)
	classHandlerAPI := _classHandler.New(classService)

	statusRepo := _statusData.New(db)
	statusService := _statusService.New(statusRepo)
	statusHandlerAPI := _statusHandler.New(statusService)

	//User
	e.POST("/login", userHandlerAPI.Login)

	// Class
	e.GET("/classes", classHandlerAPI.GetAllClass, middlewares.JWTMiddleware())
	e.POST("/classes", classHandlerAPI.CreateClass, middlewares.JWTMiddleware())
	e.GET("/classes/:class_id", classHandlerAPI.GetClassById, middlewares.JWTMiddleware())
	e.PUT("/classes/:class_id", classHandlerAPI.UpdateClass, middlewares.JWTMiddleware())
	e.DELETE("/classes/:class_id", classHandlerAPI.DeleteClass, middlewares.JWTMiddleware())

	// Status
	e.GET("/statuses", statusHandlerAPI.GetAllStatus, middlewares.JWTMiddleware())

}
