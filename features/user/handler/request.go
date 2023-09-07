package handler

import "kelompok1/immersive-dash/features/user"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	TeamId   uint   `json:"team_id"`
	Status   string `json:"status"`
}

func RequestToCore(input UserRequest) user.CoreUser {
	return user.CoreUser{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		TeamId:   input.TeamId,
		Status:   input.Status,
	}
}
func MapReqToCoreUser(req UserRequest) user.CoreUser {
	return user.CoreUser{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
		TeamId:   req.TeamId,
		Status:   req.Status,
	}

}
