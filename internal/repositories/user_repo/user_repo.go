package user_repo

import "gorm.io/gorm"

type UserRepo interface {
	GetUser(userID int64) (*User, error)
	GetUsers(initiatorUserID int64, userFilter *UserFilter) ([]User, error)
	GetUserCredentials(userID int64) (*UserCredentials, error)
	CreateUser(user User) (*User, error)
	CreateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error)
	UpdateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error)
}

type UserRepoPG struct {
	db *gorm.DB
}

func (u UserRepoPG) GetUser(userID int64) (*User, error) {
	user := User{}
	res := u.db.Model(User{}).First(&user, userID)
	return &user, res.Error
}

func (u UserRepoPG) GetUsers(initiatorUserID int64, userFilter *UserFilter) ([]User, error) {
	panic("implement me")
}

func (u UserRepoPG) GetUserCredentials(userID int64) (*UserCredentials, error) {
	userCredentials := UserCredentials{}
	res := u.db.Model(UserCredentials{}).First(&userCredentials, userID)
	return &userCredentials, res.Error
}

func (u UserRepoPG) CreateUser(user User) (*User, error) {
	res := u.db.Model(User{}).Create(&user)
	return &user, res.Error
}

func (u UserRepoPG) CreateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error) {
	res := u.db.Model(UserCredentials{}).Create(&userCredentials)
	return &userCredentials, res.Error
}

func (u UserRepoPG) UpdateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error) {
	res := u.db.Model(UserCredentials{}).Save(&userCredentials)
	return &userCredentials, res.Error
}

func NewUserRepo(db *gorm.DB) *UserRepoPG {
	return &UserRepoPG{
		db: db,
	}
}
