package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"net/http"
	"strconv"
)

func RemoveUserFromRoomHandler(log configs.Logger, roomSvc chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomIDstr := c.Params("room_id")
		roomID, err := strconv.Atoi(roomIDstr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		userIDstr := c.Params("user_id")
		userID, err := strconv.Atoi(userIDstr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		if _, err := roomSvc.DeleteUserFromRoom(int64(userID), int64(roomID)); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK).Send(nil)
		return nil
	}
}
