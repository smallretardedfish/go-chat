package user_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"github.com/smallretardedfish/go-chat/tools/slice"
	"net/http"
)

func GetUsersHandler(log logging.Logger, service user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		filter := &user.UserFilter{}
		searchStr := c.Query("search")

		if searchStr != "" {
			filter.Search = &searchStr
		}
		domainUsers, err := service.GetUsers(filter)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		users := slice.Map(domainUsers, domainUserToUser)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		return c.Status(http.StatusOK).JSON(users)
	}
}
