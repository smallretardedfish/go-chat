package user_cred_repo

import (
	"errors"
	"gorm.io/gorm"
)

type UserCredentialsRepo interface {
	CreateUserCredentials(credentials UserCredentials) (*UserCredentials, error)
	GetUserCredentials(email string) (*UserCredentials, error)
}

type UserCredentialsPG struct {
	db *gorm.DB
}

func (u *UserCredentialsPG) CreateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error) {
	err := u.db.Create(&userCredentials).Error
	if err != nil {
		return nil, err
	}
	return &userCredentials, nil
}

func (u *UserCredentialsPG) GetUserCredentials(email string) (*UserCredentials, error) {
	userCredentials := UserCredentials{}
	err := u.db.Model(UserCredentials{}).First(&userCredentials, email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userCredentials, nil
}

func (u *UserCredentialsPG) UpdateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error) {
	err := u.db.Model(UserCredentials{}).Save(&userCredentials).Error
	if err != nil {
		return nil, err
	}
	return &userCredentials, nil
}

func NewUserCredentialsPG(db *gorm.DB) *UserCredentialsPG {
	return &UserCredentialsPG{db: db}
}
