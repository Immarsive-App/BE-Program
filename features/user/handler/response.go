package handler

import "kelompok1/immersive-dash/features/user"

type LoginResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}
type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// mapping from userCore to LoginResponse
func CoreToLoginRes(Core user.CoreUser, jwtToken string) LoginResponse {
	return LoginResponse{
		Id:    Core.ID,
		Email: Core.Email,
		Token: jwtToken,
		Role:  Core.Role,
	}
}
