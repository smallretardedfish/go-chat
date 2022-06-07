package user

import (
	"errors"
	"fmt"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"golang.org/x/crypto/bcrypt"
)

var EmailDoesNotExist = errors.New("email not present")

type AuthService interface {
	SingIn(email, password string) (*User, error) // bcrypt
	SignUp(user User, credentials UserCredentials) (*User, error)
	UpdatePassword(userID int64, password string) (bool, error)
}

type AuthServiceImpl struct {
	userCredRepo user_cred_repo.UserCredentialsRepo
	userRepo     user_repo.UserRepo
}

func NewAuthServiceImpl(userCredRepo user_cred_repo.UserCredentialsRepo, userRepo user_repo.UserRepo) *AuthServiceImpl {
	return &AuthServiceImpl{userCredRepo: userCredRepo, userRepo: userRepo}
}

func (a *AuthServiceImpl) SingIn(email, password string) (*User, error) {
	credentials, err := a.userCredRepo.GetUserCredentials(email)
	if err != nil {
		return nil, fmt.Errorf("error while looking for email: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password)); err != nil {
		return nil, err
	}
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	domainUser := repoUserToDomainUser(*user)

	return &domainUser, nil
}

func (a *AuthServiceImpl) SignUp(user User, credentials UserCredentials) (*User, error) {
	repoCreds := domainCredentialsToRepoCredentials(credentials)
	_, err := a.userCredRepo.CreateUserCredentials(repoCreds)
	if err != nil {
		return nil, err
	}
	domainUser := domainUserToRepoUser(user)
	createdUser, err := a.userRepo.CreateUser(domainUser)
	if err != nil {
		return nil, err
	}
	usr := repoUserToDomainUser(*createdUser)
	return &usr, nil
}

func (a *AuthServiceImpl) UpdatePassword(id int64, oldPassword, newPassword string) (bool, error) { // why id ???
	//TODO implement method
	panic("implement me")
}
