package auth_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
	"time"
)

var jwtKey = []byte("jwt-secret") //  TODO check how to store this thing

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func SignInHandler(log configs.Logger, authSvc user.AuthService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authCredentials := UserCredentials{}
		if err := c.BodyParser(&authCredentials); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		log.Info(authCredentials)
		signedUser, err := authSvc.SingIn(authCredentials.Email, authCredentials.Password)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if signedUser == nil {
			c.Status(http.StatusUnauthorized)
			return nil
		}
		//TODO change JWT token expiration time
		expTime := time.Now().Add(30 * 24 * time.Hour).Unix() // 30 days just for local testing
		claims := Claims{
			UserID: int(signedUser.ID),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTime,
			},
		}

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := jwtToken.SignedString(jwtKey)
		if err != nil {
			log.Warn(err)
			c.Status(http.StatusInternalServerError)
			return err
		}

		c.Set("Token", tokenString)

		return nil
	}
}
