package data

import (
	"errors"
	"fmt"
	"kelompok1/immersive-dash/features/user"
	"kelompok1/immersive-dash/helpers"

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

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.CoreUser) error {
	// mapping dari struct core to struct gorm model
	userGorm := User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		TeamId:   input.TeamId,
		Role:     input.Role,
		//Status:   input.Status,
	}
	//userGorm.Password = helpers.HashPassword(input.Password)
	// simpan ke DB
	tx := repo.db.Create(&userGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
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

// SelectAll implements user.UserDataInterface.
func (repo *userQuery) SelectAll() ([]user.CoreUser, error) {
	var usersData []User
	tx := repo.db.Find(&usersData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println("users:", usersData)
	//mapping dari struct gorm model ke struct CoreUser (entity)
	var usersCoreUser []user.CoreUser
	for _, value := range usersData {
		var user = user.CoreUser{
			ID:       value.ID,
			FullName: value.FullName,
			Email:    value.Email,
			Password: value.Password,
			TeamId:   value.TeamId,
			Role:     value.Role,
			//Status:    value.Status,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		usersCoreUser = append(usersCoreUser, user)
	}
	return usersCoreUser, nil
}

func (repo *userQuery) GetByID(id uint) (user.CoreUser, error) {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {

			return user.CoreUser{}, fmt.Errorf("User with ID %d not found", id)
		}
		return user.CoreUser{}, tx.Error
	}

	userCoreUser := user.CoreUser{
		ID:       userGorm.ID,
		FullName: userGorm.FullName,
		Email:    userGorm.Email,
		Password: userGorm.Password,
		TeamId:   userGorm.TeamId,
		Role:     userGorm.Role,
		//Status:    userGorm.Status,
		CreatedAt: userGorm.CreatedAt,
		UpdatedAt: userGorm.UpdatedAt,
	}

	return userCoreUser, nil
}

func (repo *userQuery) Delete(id uint) error {
	var userGorm User
	tx := repo.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *userQuery) Update(id uint, updatedUser user.CoreUser) error {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("User with ID %d not found", id)
		}
		return tx.Error
	}

	userGorm.FullName = updatedUser.FullName
	userGorm.Email = updatedUser.Email
	userGorm.Password = updatedUser.Password
	userGorm.TeamId = updatedUser.TeamId
	userGorm.Role = updatedUser.Role
	//userGorm.Status = updatedUser.Status

	tx = repo.db.Save(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if updatedUser.Password != "" {
		updatedUser.Password = helpers.HashPassword(updatedUser.Password)

	}
	return nil

}
