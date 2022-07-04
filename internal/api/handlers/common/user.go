package common

import (
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
)

type User struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func DomainUserToUser(u user.User) User {
	return User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func UserToDomainUser(u User) user.User {
	return user.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func UserToChatUser(u User) chat.User {
	return chat.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

func ChatUserToUser(u chat.User) User {
	return User{
		ID:   u.ID,
		Name: u.Name,
	}
}
