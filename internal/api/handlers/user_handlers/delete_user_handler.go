package user_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
)

//TODO reseach error handler
func DeleteUserHandler(log logging.Logger, service user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		currentUser := c.Context().UserValue("user").(*user.User)

		_, err := service.DeleteUser(currentUser.ID) //TODO how to use this boolean
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}
}
