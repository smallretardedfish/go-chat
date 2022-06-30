package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"github.com/smallretardedfish/go-chat/tools/slice"
	"net/http"
	"strconv"
)

func GetRoomsHandler(log logging.Logger, service chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		limitStr := c.Query("limit", "10")
		offsetStr := c.Query("offset", "0")
		limit, err := strconv.Atoi(limitStr) //TODO change Atoi to ParseInt
		offset, err := strconv.Atoi(offsetStr)
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
