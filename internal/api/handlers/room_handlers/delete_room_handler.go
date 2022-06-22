package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
	"strconv"
)

//TODO reseach error handler
func DeleteRoomHandler(log configs.Logger, service chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomIdStr := c.Params("id")
		roomID, err := strconv.Atoi(roomIdStr)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		usr := c.Context().UserValue("user").(*user.User)

		ok, err := service.DeleteRoom(usr.ID, int64(roomID))
		if err != nil {
			log.Error(err)
			c.Status(http.StatusInternalServerError)
			return err
		}
		if !ok {
			c.Status(http.StatusNotFound)
			return nil
		}
		c.Status(http.StatusOK)
		return nil
	}
}
