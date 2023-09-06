package handler

import (
	"kelompok1/immersive-dash/app/middlewares"
	"kelompok1/immersive-dash/features/user"
	"kelompok1/immersive-dash/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
func (handler *UserHandler) GetAllUser(c echo.Context) error {
	result, err := handler.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct  to struct response
	var usersResponse []UserResponse
	for _, value := range result {
		usersResponse = append(usersResponse, UserResponse{
			ID:       value.ID,
			FullName: value.FullName,
			Email:    value.Email,
			Role:     value.Role,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", usersResponse))
	// return c.JSON(http.StatusOK, map[string]any{
	// 	"code":    200,
	// 	"message": "success read data",
	// 	"data":    usersResponse,
	// })

}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	//userRole := middlewares.ExtractTokenUserRole(c)
	//if userRole != "super admin" {
	//return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "Access denied. Admin privileges required.", nil))
	//}
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	userCore := user.CoreUser{
		FullName: userInput.FullName,
		Email:    userInput.Email,
		Password: userInput.Password,
		Role:     userInput.Role,
		TeamId:   userInput.TeamId,
	}
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success insert data", nil))
}

func (h *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	dataLogin, token, err := h.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	response := CoreToLoginRes(dataLogin, token)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "login successfully", response))
}
func (handler *UserHandler) FindUserByID(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := handler.userService.GetByID(uint(id))
	if err != nil {
		if err.Error() == "Tidak ada" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "User tidak ditemukan", nil))
		}

		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	userResponse := UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Berhasil", userResponse))
}

func (handler *UserHandler) DeleteUserByID(c echo.Context) error {
	userRole := middlewares.ExtractTokenUserRole(c)
	if userRole != "super admin" {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "Access denied. Admin privileges required.", nil))
	}
	id := middlewares.ExtractTokenUserId(c)

	err := handler.userService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User berhasil terhapus", nil))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	userRole := middlewares.ExtractTokenUserRole(c)
	if userRole != "super admin" {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "Access denied. Admin privileges required.", nil))
	}
	id := middlewares.ExtractTokenUserId(c)

	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	updatedUser := user.CoreUser{
		FullName: userInput.FullName,
		Email:    userInput.Email,
		Password: userInput.Password,
		TeamId:   userInput.TeamId,
		Role:     userInput.Role,
	}

	err := handler.userService.Update(uint(id), updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating user", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User updated successfully", nil))
}
