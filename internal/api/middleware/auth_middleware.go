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
		var tokenString string
		if c.Get("Token") != "" {
			tokenString = c.Get("Token")
		} else if c.Query("token") != "" {
			tokenString = c.Query("token")
		}

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
			return c.Send(nil)
		}
		log.Info("PARSING TOKEN")
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			id := claims.UserID
			usr, err := userService.GetUser(id)
			log.Info("User got from token: ", usr)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return err
			}
			c.Context().SetUserValue("user", usr)
		}
		return c.Next()
	}
}
