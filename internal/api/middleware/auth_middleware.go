package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func AuthMiddleware(log configs.Logger, userService user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Token")
		if tokenString == "" {
			c.Status(http.StatusUnauthorized)
			c.Send(nil)
			return nil
		}
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("jwt-secret"), nil // TODO maybe retrieve secret from config
		})
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Send(nil)
			return err
		}
		log.Info("PARSING TOKEN")
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			id := claims.UserID
			getUser, err := userService.GetUser(id)
			log.Info("User got from token: ", getUser)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return err
			}
			c.Context().SetUserValue("user", getUser)
		}
		return c.Next()
	}
}
