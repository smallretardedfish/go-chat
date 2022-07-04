package user

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
)

type AuthService interface {
	SingIn(email, password string) (*User, error)
	SignUp(user User, credentials UserCredentials) (*User, error)
	UpdatePassword(userID int64, oldPassword, newPassword string) (bool, error)
}

type UserRepo interface {
	GetUserByID(userID int64) (*user_repo.User, error)
	GetUserByEmail(email string) (*user_repo.User, error)
	CreateUser(user user_repo.User) (*user_repo.User, error)
}

type UserCredentialsRepo interface {
	CreateUserCredentials(credentials user_cred_repo.UserCredentials) (*user_cred_repo.UserCredentials, error)
	GetUserCredentials(email string) (*user_cred_repo.UserCredentials, error)
	UpdateUserCredentials(credentials user_cred_repo.UserCredentials) (*user_cred_repo.UserCredentials, error)
}

type AuthServiceImpl struct {
	userCredRepo UserCredentialsRepo
	userRepo     UserRepo
	crypto       Crypto
}

type Crypto interface {
	HashAndSalt(password []byte) (string, error)
	ComparePasswords(hashedPassword string, password []byte) error
}

func NewAuthServiceImpl(userCredRepo UserCredentialsRepo, userRepo UserRepo, crypto Crypto) *AuthServiceImpl {
	return &AuthServiceImpl{userCredRepo: userCredRepo, userRepo: userRepo, crypto: crypto}
}

func (a *AuthServiceImpl) SingIn(email, password string) (*User, error) {
	credentials, err := a.userCredRepo.GetUserCredentials(email)
	if err != nil || credentials == nil {
		return nil, err
	}
	if err := a.crypto.ComparePasswords(credentials.Password, []byte(password)); err != nil {
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
	repoUser := userToRepoUser(user)
	createdUser, err := a.userRepo.CreateUser(repoUser)
	if err != nil || createdUser == nil {
		return nil, err
	}
	repoCreds := domainCredentialsToRepoCredentials(credentials)
	hashedPassword, err := a.crypto.HashAndSalt([]byte(repoCreds.Password))
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
	if err := a.crypto.ComparePasswords(credentials.Password, []byte(oldPassword)); err != nil {
		return false, nil
	}
	credentials.Password = newPassword

	if _, err := a.userCredRepo.UpdateUserCredentials(*credentials); err != nil {
		return false, err
	}
	return true, nil
}
