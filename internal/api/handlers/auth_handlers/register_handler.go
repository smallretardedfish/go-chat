package auth_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
)

func RegisterHandler(log configs.Logger, authSvc user.AuthService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userInp := &SignUpInput{}
		err := c.BodyParser(userInp)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		usr := User{
			Name:  userInp.Name,
			Email: userInp.Email,
		}
		userCreds := UserCredentials{
			Email:    userInp.Email,
			Password: userInp.Password,
		}
		domainUser := userToDomainUser(usr)
		domainUserCreds := userCredentialsToDomainUserCredentials(userCreds)

		if _, err := authSvc.SignUp(domainUser, domainUserCreds); err != nil {
			log.Error("error while signing up new user:", err)
			c.Status(http.StatusInternalServerError)
			return err
		}
		return nil
	}
}
