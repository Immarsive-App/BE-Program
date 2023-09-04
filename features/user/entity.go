package user

import "time"

type CoreUser struct {
	ID        uint
	FullName  string
	Email     string
	Password  string
	TeamId    uint
	Role      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserDataInterface interface {
}

type UserServiceInterface interface {
}
