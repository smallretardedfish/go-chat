package user

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
)

func repoCredentialsToDomainCredentials(userCredentials user_cred_repo.UserCredentials) UserCredentials {
	return UserCredentials{
		Email:    userCredentials.Email,
		Password: userCredentials.Password,
	}
}
func domainCredentialsToRepoCredentials(userCredentials UserCredentials) user_cred_repo.UserCredentials {
	return user_cred_repo.UserCredentials{
		Email:    userCredentials.Email,
		Password: userCredentials.Password,
	}
}
