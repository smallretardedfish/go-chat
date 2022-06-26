package auth_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
)

//TODO create also add refresh token

var jwtKey = []byte("jwt-secret") //  TODO move to env file

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func SignInHandler(log logging.Logger, authSvc user.AuthService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authCredentials := UserCredentials{}
		if err := c.BodyParser(&authCredentials); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		signedUser, err := authSvc.SingIn(authCredentials.Email, authCredentials.Password)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if signedUser == nil {
			c.Status(http.StatusUnauthorized)
			return nil
		}

		usr := domainUserToUser(*signedUser)
		token, err := CreateToken(usr.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		return c.Status(http.StatusOK).
			JSON(struct {
				User  User   `json:"user"`
				Token string `json:"token"`
			}{
				User:  usr,
				Token: token,
			})
	}
}
