package user_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/api/handlers/common"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
	"strconv"
)

func GetUserHandler(log logging.Logger, service user.UserService) func(c *fiber.Ctx) error {
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
		if domainUser == nil {
			c.Status(http.StatusNotFound)
			log.Warn("No such user present in repo")
			return c.JSON(nil) // probably that way
		}
		usr := common.DomainUserToUser(*domainUser)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		return c.Status(http.StatusOK).JSON(usr)
	}
}
