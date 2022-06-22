package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
	"strconv"
)

func RemoveCurrentUserFromRoomHandler(log configs.Logger, roomSvc chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomIdStr := c.Params("room_id")
		roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		initiator := c.Context().UserValue("user").(*user.User)

		if _, err := roomSvc.DeleteCurrentUser(initiator.ID, roomId); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		return c.Status(http.StatusOK).Send(nil)
	}
}
