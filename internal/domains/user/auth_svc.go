package user

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SingIn(email, password string) (*User, error)
	SignUp(user User, credentials UserCredentials) (*User, error)
	UpdatePassword(userID int64, oldPassword, newPassword string) (bool, error)
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
	if err != nil || credentials == nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password)); err != nil {
		return nil, nil
	}
	repoUser, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	user := repoUserToUser(*repoUser)

	return &user, nil
}

func (a *AuthServiceImpl) SignUp(user User, credentials UserCredentials) (*User, error) {
	domainUser := userToRepoUser(user)
	createdUser, err := a.userRepo.CreateUser(domainUser)
	if err != nil || createdUser == nil {
		return nil, err
	}
	repoCreds := domainCredentialsToRepoCredentials(credentials)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(repoCreds.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	repoCreds.Password = string(hashedPassword)

	if _, err := a.userCredRepo.CreateUserCredentials(repoCreds); err != nil {
		return nil, err
	}
	usr := repoUserToUser(*createdUser)
	return &usr, nil
}

func (a *AuthServiceImpl) UpdatePassword(userID int64, oldPassword, newPassword string) (bool, error) {

	usr, err := a.userRepo.GetUserByID(userID)
	if err != nil || usr == nil {
		return false, err
	}
	credentials, err := a.userCredRepo.GetUserCredentials(usr.Email)
	if err != nil {
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(oldPassword)); err != nil {
		return false, nil
	}
	credentials.Password = newPassword

	if _, err := a.userCredRepo.UpdateUserCredentials(*credentials); err != nil {
		return false, err
	}
	return true, nil
}
