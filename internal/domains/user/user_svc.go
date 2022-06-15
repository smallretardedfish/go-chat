package user

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"github.com/smallretardedfish/go-chat/tools/slice"
)

type UserService interface {
	GetUser(userID int64) (*User, error)
	GetUsers(userFilter *UserFilter) ([]User, error)
	UpdateUser(user User) (*User, error)
	DeleteUser(userID int64) (bool, error)
}

type UserServiceImpl struct {
	userRepo user_repo.UserRepo
}

func (u *UserServiceImpl) GetUsers(userFilter *UserFilter) ([]User, error) {
	userFil := (*user_repo.UserFilter)(userFilter) // tricky thing, converting because of same structure
	users, err := u.userRepo.GetUsers(userFil)
	if err != nil {
		return nil, err
	}
	return slice.Map(users, repoUserToUser), nil
}

func (u *UserServiceImpl) GetUser(userID int64) (*User, error) {
	usr, err := u.userRepo.GetUserByID(userID)
	if err != nil || usr == nil {
		return nil, err
	}
	user := repoUserToUser(*usr)
	return &user, nil
}

func (u *UserServiceImpl) UpdateUser(user User) (*User, error) {
	userToUpdate := userToRepoUser(user)
	usr, err := u.userRepo.UpdateUser(userToUpdate)
	if err != nil {
		return nil, err
	}
	user = repoUserToUser(*usr)
	return &user, nil
}

func (u *UserServiceImpl) DeleteUser(userID int64) (bool, error) {
	ok, err := u.userRepo.DeleteUser(userID)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func NewUserServiceImpl(userRepo user_repo.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo}
}
