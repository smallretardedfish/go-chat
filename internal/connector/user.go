package connector

import "github.com/smallretardedfish/go-chat/internal/domains/user"

type User struct {
	ID       int64
	Username string
}

func DomainUserToUser(domainUser user.User) *User {
	return &User{
		ID:       domainUser.ID,
		Username: domainUser.Name,
	}
}
