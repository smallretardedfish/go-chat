package user_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
	"strconv"
)

//TODO reseach error handler
func DeleteUserHandler(log configs.Logger, service user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userIdStr := c.Params("id")
		id, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		_, err = service.DeleteUser(int64(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}
}
