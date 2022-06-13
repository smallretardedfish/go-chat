package auth_handlers

import (
	"github.com/smallretardedfish/go-chat/internal/domains/user"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func userToDomainUser(usr User) user.User {
	return user.User{
		ID:    usr.ID,
		Name:  usr.Name,
		Email: usr.Email,
	}
}

func domainUserToUser(usr user.User) User {
	return User{
		Name:  usr.Name,
		Email: usr.Email,
	}
}

type UserCredentials struct {
	Email    string
	Password string
}

func domainUserCredentialsToUserCredentials(usr user.UserCredentials) UserCredentials {
	return UserCredentials{
		Email:    usr.Email,
		Password: usr.Password,
	}
}

func userCredentialsToDomainUserCredentials(usrCreds UserCredentials) user.UserCredentials {
	return user.UserCredentials{
		Email:    usrCreds.Email,
		Password: usrCreds.Password,
	}
}
