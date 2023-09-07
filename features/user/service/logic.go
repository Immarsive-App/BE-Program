package service

import (
	"errors"
	"kelompok1/immersive-dash/app/middlewares"
	"kelompok1/immersive-dash/features/user"

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

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.CoreUser) error {
	if input.FullName == "" || input.Email == "" || input.Password == "" {
		return errors.New("validation error. name/email/password required")
	}
	err := service.userData.Insert(input)
	return err
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
	// checkPassword := helpers.ComparePassword(password, dataLogin.Password)
	// if !checkPassword {
	// 	return user.CoreUser{}, "", errors.New("login failed, wrong password")
	// }
	token, err := middlewares.CreateToken(int(dataLogin.ID))
	if err != nil {
		log.Error("Create token error:", err)
		return user.CoreUser{}, "", err
	}
	return dataLogin, token, nil
}

// GetAll implements user.UserServiceInterface.
func (service *userService) GetAll() ([]user.CoreUser, error) {
	result, err := service.userData.SelectAll()
	return result, err
}

func (service *userService) GetByID(id uint) (user.CoreUser, error) {
	if id == 0 {
		return user.CoreUser{}, errors.New("ID tidak valid")
	}
	result, err := service.userData.GetByID(id)
	return result, err
}

func (service *userService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validation error. invalid ID")
	}
	err := service.userData.Delete(id)
	return err
}

func (service *userService) Update(id uint, updatedUser user.CoreUser) error {
	if id == 0 {
		return errors.New("validation error. invalid ID")
	}
	err := service.userData.Update(id, updatedUser)
	return err
}
