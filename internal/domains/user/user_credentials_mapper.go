package user

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
)

func repoUserCredentialsToUserCredentials(userCredentialsDAO user_cred_repo.UserCredentials) UserCredentials {
	return UserCredentials{
		Email:    userCredentialsDAO.Email,
		Password: userCredentialsDAO.Password,
	}
}
