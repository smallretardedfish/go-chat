package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
	"strconv"
)

func GetRoomHandler(log logging.Logger, service chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomIdStr := c.Params("id")
		roomID, err := strconv.ParseInt(roomIdStr, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		usr := c.Context().UserValue("user").(*user.User)

		domainRoom, err := service.GetRoom(usr.ID, roomID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if domainRoom == nil {
			c.Status(http.StatusNotFound)
			return nil
		}
		room := domainRoomToRoom(*domainRoom)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		return c.Status(http.StatusOK).JSON(room)
	}
}
