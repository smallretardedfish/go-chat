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

		signedUser, err := authSvc.SignUp(domainUser, domainUserCreds)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if signedUser == nil {
			c.Status(http.StatusUnauthorized)
			return nil
		}
		token, err := CreateToken(signedUser.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		return c.Status(http.StatusOK).
			JSON(struct {
				User  User   `json:"user"`
				Token string `json:"token"`
			}{
				User:  domainUserToUser(*signedUser),
				Token: token,
			})
	}
}
