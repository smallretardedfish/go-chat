package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"github.com/smallretardedfish/go-chat/pkg/slice"
	"net/http"
	"strconv"
)

func GetRoomsHandler(log logging.Logger, service chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		limitStr := c.Query("limit", "10")
		offsetStr := c.Query("offset", "0")
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		offset, err := strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		usr := c.Context().UserValue("user").(*user.User)

		domainRooms, err := service.GetRooms(int64(limit), int64(offset), usr.ID)
		rooms := slice.Map(domainRooms, domainRoomToRoom)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return c.JSON(rooms)
	}
}
