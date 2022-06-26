package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
)

func RemoveUsersFromRoomHandler(log logging.Logger, roomSvc chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		removeData := &struct {
			RoomID   int64   `json:"room_id"`
			ToRemove []int64 `json:"to_remove"`
		}{}

		initiator := c.Context().UserValue("user").(*user.User)

		if _, err := roomSvc.DeleteUsersFromRoom(initiator.ID, removeData.RoomID, removeData.ToRemove); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		return c.Status(http.StatusOK).Send(nil)
	}
}
