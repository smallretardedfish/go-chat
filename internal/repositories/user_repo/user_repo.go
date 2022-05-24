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

type UserRepoImpl struct {
	db *gorm.DB
}

func (u UserRepoImpl) GetUser(userID int64) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoImpl) GetUsers(initiatorUserID int64, userFilter *UserFilter) ([]User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoImpl) GetUserCredentials(userID int64) (*UserCredentials, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoImpl) CreateUser(user User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoImpl) CreateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoImpl) UpdateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{
		db: db,
	}
}
