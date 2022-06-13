package auth_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
)

func SignInHandler(log configs.Logger, service user.AuthService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Body()
		return nil
	}
}
