package user

import "github.com/smallretardedfish/go-chat/internal/repositories/user_repo"

type UserService interface {
	GetUser(userID int64) (*User, error)
	UpdateUser(user User) (*User, error)
	DeleteUser(userID int64) (bool, error)
}

type UserServiceImpl struct {
	userRepo user_repo.UserRepo
}

func (u *UserServiceImpl) GetUser(userID int64) (*User, error) {
	usr, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	user := repoUserToDomainUser(*usr)
	return &user, nil
}

func (u *UserServiceImpl) UpdateUser(user User) (*User, error) {
	userToUpdate := domainUserToRepoUser(user)
	usr, err := u.userRepo.UpdateUser(userToUpdate)
	if err != nil {
		return nil, err
	}
	user = repoUserToDomainUser(*usr)
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
