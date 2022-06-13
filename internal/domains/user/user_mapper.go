package user

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
)

func repoUserToUser(user user_repo.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func userToRepoUser(user User) user_repo.User {
	return user_repo.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
