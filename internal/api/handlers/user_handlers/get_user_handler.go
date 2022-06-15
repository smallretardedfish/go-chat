package user_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
	"strconv"
)

//TODO reseach error handler
func GetUserHandler(log configs.Logger, service user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userIdStr := c.Params("id")
		id, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		domainUser, err := service.GetUser(int64(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if domainUser == nil { // TODO research how to handle no user
			log.Warn("No such user present in repo")
			c.JSON(nil) // probably that way
			return nil
		}
		usr := domainUserToUser(*domainUser)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		return c.Status(http.StatusOK).JSON(usr)
	}
}
