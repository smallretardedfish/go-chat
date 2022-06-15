package user_handlers

import "github.com/smallretardedfish/go-chat/internal/domains/user"

type User struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func domainUserToUser(u user.User) User {
	return User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
