package room_handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
	"strconv"
)

//TODO reseach error handler
func GetRoomHandler(log configs.Logger, service chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomIdStr := c.Params("id")
		roomID, err := strconv.Atoi(roomIdStr)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		log.Info(c.Context().UserValue("user"))
		usr, ok := c.Context().UserValue("user").(*user.User)
		log.Info(usr)
		if !ok {
			return fmt.Errorf("error: can`t assert user from context to *user.User")
		}
		domainRoom, err := service.GetRoom(usr.ID, int64(roomID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if domainRoom == nil { // TODO research how to handle if user got no room
			c.JSON(nil) // probably that way
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
