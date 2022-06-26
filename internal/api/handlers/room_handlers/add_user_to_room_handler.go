package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
)

func AddUserToRoomHandler(log logging.Logger, roomSvc chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomUser := &struct {
			RoomID  int64   `json:"room_id"`
			UserIDs []int64 `json:"user_id"`
		}{}

		if err := c.BodyParser(roomUser); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if _, err := roomSvc.AddUsersToRoom(roomUser.UserIDs, roomUser.RoomID); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		return c.Status(http.StatusOK).Send(nil)
	}
}
