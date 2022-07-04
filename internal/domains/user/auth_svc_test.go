package user

import (
	mock "github.com/golang/mock/gomock"
	"github.com/smallretardedfish/go-chat/internal/domains/user/mocks"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"github.com/smallretardedfish/go-chat/pkg/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuthServiceImpl_SingUp(t *testing.T) {
	ctrl := mock.NewController(t)
	defer ctrl.Finish()

	realCrypto := &crypto.AppCrypto{}

	userRepoMock := mocks.NewMockUserRepo(ctrl)
	userCredRepoMock := mocks.NewMockUserCredentialsRepo(ctrl)
	cryptoMock := mocks.NewMockCrypto(ctrl)

	authSvc := NewAuthServiceImpl(userCredRepoMock, userRepoMock, cryptoMock)

	name := "BRUH"
	email := "bruh2002@mail.ua"
	password := "zxc123"
	hashedPassword, _ := realCrypto.HashAndSalt([]byte(password))

	usr := User{
		Name:  name,
		Email: email,
	}
	usrCreds := UserCredentials{
		Email:    email,
		Password: password,
	}
	userRepoMock.EXPECT().CreateUser(user_repo.User{
		Name:  name,
		Email: email,
	}).Return(&user_repo.User{
		Name:  name,
		Email: email,
	}, nil)

	cryptoMock.EXPECT().HashAndSalt([]byte(password)).Return(hashedPassword, nil)

	userCredRepoMock.EXPECT().CreateUserCredentials(user_cred_repo.UserCredentials{
		Email:    email,
		Password: hashedPassword,
	}).Return(&user_cred_repo.UserCredentials{
		Email:    email,
		Password: hashedPassword,
	}, nil)

	_, err := authSvc.SignUp(usr, usrCreds)
	require.NoError(t, err)
}

func TestAuthServiceImpl_SingIn(t *testing.T) {
	ctrl := mock.NewController(t)
	defer ctrl.Finish()

	realCrypto := &crypto.AppCrypto{}

	userRepoMock := mocks.NewMockUserRepo(ctrl)
	userCredRepoMock := mocks.NewMockUserCredentialsRepo(ctrl)
	cryptoMock := mocks.NewMockCrypto(ctrl)

	authSvc := NewAuthServiceImpl(userCredRepoMock, userRepoMock, cryptoMock)

	name := "BRUH"
	email := "bruh2002@mail.ua"
	password := "zxc123"
	hashedPassword, _ := realCrypto.HashAndSalt([]byte(password))

	userRepoMock.EXPECT().GetUserByEmail(email).Return(&user_repo.User{
		Name:  name,
		Email: email,
	}, nil)

	userCredRepoMock.EXPECT().GetUserCredentials(email).Return(&user_cred_repo.UserCredentials{
		Email:    email,
		Password: hashedPassword,
	}, nil)

	cryptoMock.EXPECT().ComparePasswords(hashedPassword, []byte(password)).Return(nil)
	_, err := authSvc.SingIn(email, password)
	require.NoError(t, err)
}
