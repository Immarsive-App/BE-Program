package data

import (
	"errors"
	"kelompok1/immersive-dash/features/user"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Login implements user.UserDataInterface.
func (r *userQuery) Login(email string) (user.CoreUser, error) {
	var dataUser User
	tx := r.db.Where("email=?", email).First(&dataUser)
	if tx.Error != nil {
		log.Error("Database error:", tx.Error)
		return user.CoreUser{}, errors.New(tx.Error.Error() + ", invalid email")
	}

	dataCore := ModelToCore(dataUser)
	return dataCore, nil
}
