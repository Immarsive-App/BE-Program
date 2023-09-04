package handler

import (
	"kelompok1/immersive-dash/features/class"
)

type ClassHandler struct {
	classService class.ClassServiceInterface
}

func New(service class.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classService: service,
	}
}
