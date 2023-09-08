package handler

import (
	"kelompok1/immersive-dash/app/middlewares"
	"kelompok1/immersive-dash/features/user"
	"kelompok1/immersive-dash/helpers"
	"net/http"
	"strconv"
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
	users, err := handler.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error fetching users", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success", users))
}
func (handler *UserHandler) CreateUser(c echo.Context) error {
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
	id := c.Param("user_id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid user ID", nil))
	}

	user, err := handler.userService.GetByID(uint(userID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "user not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error fetching user", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success", user))
}

func (handler *UserHandler) DeleteUserByID(c echo.Context) error {
	id := c.Param("user_id")

	userRole := middlewares.ExtractTokenUserRole(c)
	if userRole == "user" {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "Access denied. Admin privileges required.", nil))
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid user ID", nil))
	}

	if err := handler.userService.Delete(uint(userID)); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "user not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error deleting user", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success deleting user", nil))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("user_id")

	userRole := middlewares.ExtractTokenUserRole(c)
	if userRole == "user" {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "Access denied. Admin privileges required.", nil))
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid user ID", nil))
	}

	// Parse and validate the updated user data from the request body.
	updatedUserInput := new(UserRequest)
	if err := c.Bind(updatedUserInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	// Create a CoreUser instance with the updated data.
	updatedUser := user.CoreUser{
		FullName: updatedUserInput.FullName,
		Email:    updatedUserInput.Email,
		Password: updatedUserInput.Password,
		Role:     updatedUserInput.Role,
		TeamId:   updatedUserInput.TeamId,
		Status:   updatedUserInput.Status,
	}

	// Call the UpdateUser method in the service layer to update the user.
	if err := handler.userService.Update(uint(userID), updatedUser); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "user not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating user", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success updating user", nil))
}
