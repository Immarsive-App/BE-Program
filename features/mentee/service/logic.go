package service

import "kelompok1/immersive-dash/features/mentee"

type menteeService struct {
	menteeData mentee.MenteeDataInterface
}

func New(repo mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
	}
}
