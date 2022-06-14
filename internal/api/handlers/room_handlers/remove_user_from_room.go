package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"net/http"
)

func RemoveUserFromRoomHandler(log configs.Logger, roomSvc chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomUser := &struct {
			RoomID int64 `json:"room_id"`
			UserID int64 `json:"user_id"`
		}{}
		err := c.BodyParser(roomUser)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if _, err := roomSvc.DeleteUserFromRoom(roomUser.UserID, roomUser.RoomID); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK).Send(nil)
		return nil
	}
}
