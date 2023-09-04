package handler

import (
	"kelompok1/immersive-dash/features/user"
	"kelompok1/immersive-dash/helpers"
	"net/http"

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
