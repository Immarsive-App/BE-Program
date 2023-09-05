package handler

import "kelompok1/immersive-dash/features/class"

type ClassRequest struct {
	ClassName string `json:"class_name"`
}

func RequestToCore(input ClassRequest) class.CoreClass {
	return class.CoreClass{
		ClassName: input.ClassName,
	}
}
