package service

import (
	"errors"
	"kelompok1/immersive-dash/app/middlewares"
	"kelompok1/immersive-dash/features/user"
	"kelompok1/immersive-dash/helpers"

	"github.com/labstack/gommon/log"
)

type userService struct {
	userData user.UserDataInterface
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

// Login implements user.UserServiceInterface.
func (s *userService) Login(email string, password string) (user.CoreUser, string, error) {
	if email == "" {
		return user.CoreUser{}, "", errors.New("email is required")
	}
	if password == "" {
		return user.CoreUser{}, "", errors.New("password is required")
	}
	dataLogin, err := s.userData.Login(email)
	if err != nil {
		log.Error("Service error:", err)
		return user.CoreUser{}, "", err
	}
	checkPassword := helpers.ComparePassword(password, dataLogin.Password)
	if !checkPassword {
		return user.CoreUser{}, "", errors.New("login failed, wrong password")
	}
	token, err := middlewares.CreateToken(int(dataLogin.ID))
	if err != nil {
		log.Error("Create token error:", err)
		return user.CoreUser{}, "", err
	}
	return dataLogin, token, nil
}
