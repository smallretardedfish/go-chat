package user_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/tools/slice"
	"net/http"
	"strconv"
)

//TODO reseach error handler
func GetUsersHandler(log configs.Logger, service user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		filter := &user.UserFilter{} // TODO probably create filter on current layer
		roomIDstr := c.Query("room")
		searchStr := c.Query("search")

		if roomIDstr == "" && searchStr == "" {
			filter = nil // no filtering in this case
		}

		if roomIDstr != "" {
			roomID, err := strconv.Atoi(roomIDstr)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return err
			}
			id := int64(roomID)
			filter.RoomID = &id
		}
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
