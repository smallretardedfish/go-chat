package user_repo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByID(userID int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUsers(userFilter *UserFilter) ([]User, error)
	CreateUser(user User) (*User, error)
	UpdateUser(user User) (*User, error)
	DeleteUser(userID int64) (bool, error)
}

type UserRepoPG struct {
	db *gorm.DB
}

var _ UserRepo = (*UserRepoPG)(nil)

func (u *UserRepoPG) GetUserByID(userID int64) (*User, error) {
	user := User{}
	err := u.db.Model(User{}).First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepoPG) GetUserByEmail(email string) (*User, error) {
	user := User{}
	err := u.db.Model(User{}).First(&user, email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepoPG) GetUsers(userFilter *UserFilter) ([]User, error) {
	var users []User

	res := u.db
	if userFilter != nil {
		if userFilter.RoomID != nil {
			res = res.Table("users AS u").
				Joins("JOIN room_users AS ru  ON u.id=ru.user_id").
				Where("room_id = ?", *userFilter.RoomID)
		}
		if userFilter.Search != nil {
			res = res.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *userFilter.Search))
		}
	}
	err := res.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepoPG) CreateUser(user User) (*User, error) {
	err := u.db.Model(User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepoPG) UpdateUser(user User) (*User, error) {
	err := u.db.Model(User{}).Where("id = ?", user.ID).Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepoPG) DeleteUser(userID int64) (bool, error) {
	err := u.db.Delete(User{}, userID).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUserRepo(db *gorm.DB) *UserRepoPG {
	return &UserRepoPG{
		db: db,
	}
}
