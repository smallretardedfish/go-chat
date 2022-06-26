package user_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
)

func UpdateUserHandler(log logging.Logger, userService user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userData := User{}
		if err := c.BodyParser(&userData); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		userToBeUpdated := userToDomainUser(userData)
		u, err := userService.UpdateUser(userToBeUpdated)
		if u == nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}
}
