package user

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
)

// TODO make mapper from repo layer to service

func repoUserToServiceUser(user user_repo.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func serviceUserToRepoUser(user User) user_repo.User {
	return user_repo.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
