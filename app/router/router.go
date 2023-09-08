package router

import (
	"kelompok1/immersive-dash/app/middlewares"
	_userData "kelompok1/immersive-dash/features/user/data"
	_userHandler "kelompok1/immersive-dash/features/user/handler"
	_userService "kelompok1/immersive-dash/features/user/service"

	_classData "kelompok1/immersive-dash/features/class/data"
	_classHandler "kelompok1/immersive-dash/features/class/handler"
	_classService "kelompok1/immersive-dash/features/class/service"

	_menteeData "kelompok1/immersive-dash/features/mentee/data"
	_menteeHandler "kelompok1/immersive-dash/features/mentee/handler"
	_menteeService "kelompok1/immersive-dash/features/mentee/service"
	_statusData "kelompok1/immersive-dash/features/status/data"
	_statusHandler "kelompok1/immersive-dash/features/status/handler"
	_statusService "kelompok1/immersive-dash/features/status/service"

	_teamData "kelompok1/immersive-dash/features/team/data"
	_teamHandler "kelompok1/immersive-dash/features/team/handler"
	_teamService "kelompok1/immersive-dash/features/team/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := _userData.New(db)
	userService := _userService.New(userRepo)
	userHandlerAPI := _userHandler.New(userService)

	classRepo := _classData.New(db)
	classService := _classService.New(classRepo)
	classHandlerAPI := _classHandler.New(classService)

	menteeRepo := _menteeData.New(db)
	menteeService := _menteeService.New(menteeRepo)
	menteeHandlerAPI := _menteeHandler.New(menteeService)
	statusRepo := _statusData.New(db)
	statusService := _statusService.New(statusRepo)
	statusHandlerAPI := _statusHandler.New(statusService)

	teamRepo := _teamData.New(db)
	teamService := _teamService.New(teamRepo)
	teamHandlerAPI := _teamHandler.New(teamService)

	//User
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/:user_id", userHandlerAPI.FindUserByID, middlewares.JWTMiddleware())
	e.POST("/users", userHandlerAPI.CreateUser)
	e.DELETE("/users/:user_id", userHandlerAPI.DeleteUserByID, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())

	// Class
	e.GET("/classes", classHandlerAPI.GetAllClass, middlewares.JWTMiddleware())
	e.POST("/classes", classHandlerAPI.CreateClass, middlewares.JWTMiddleware())
	e.GET("/classes/:class_id", classHandlerAPI.GetClassById, middlewares.JWTMiddleware())
	e.PUT("/classes/:class_id", classHandlerAPI.UpdateClass, middlewares.JWTMiddleware())
	e.DELETE("/classes/:class_id", classHandlerAPI.DeleteClass, middlewares.JWTMiddleware())

	//mentee
	e.POST("/mentees", menteeHandlerAPI.CreateMentee, middlewares.JWTMiddleware())
	e.GET("/mentees", menteeHandlerAPI.GetAllMentee, middlewares.JWTMiddleware())
	e.GET("/mentees/:mentee_id", menteeHandlerAPI.GetMenteeById, middlewares.JWTMiddleware())
	e.GET("/mentees/:mentee_id/feedbacks", menteeHandlerAPI.GetMenteeFeedback, middlewares.JWTMiddleware())
	e.PUT("/mentees/:mentee_id", menteeHandlerAPI.UpdateMenteeById, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:mentee_id", menteeHandlerAPI.DeleteMenteeById, middlewares.JWTMiddleware())
	e.GET("/mentees/:mentee_id", menteeHandlerAPI.GetMenteeById, middlewares.JWTMiddleware())
	// Status
	e.GET("/statuses", statusHandlerAPI.GetAllStatus, middlewares.JWTMiddleware())

	// Team
	e.GET("/teams", teamHandlerAPI.GetAllTeam, middlewares.JWTMiddleware())

}
