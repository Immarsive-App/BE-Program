package service

import "kelompok1/immersive-dash/features/class"

type classService struct {
	classData class.ClassDataInterface
}

func New(repo class.ClassDataInterface) class.ClassServiceInterface {
	return &classService{
		classData: repo,
	}
}
