package service

import (
	"kelompok1/immersive-dash/features/user"
)

type userService struct {
	userData user.UserDataInterface
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}
