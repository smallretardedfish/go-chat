package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
)

func AuthMiddleware(log configs.Logger, userService user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Get("token")
		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Send(nil)
			return nil
		}
		//1) check token validity
		//2) get user_id from token
		//3) get user by id from service and set to context

		return c.Next()
	}
}
