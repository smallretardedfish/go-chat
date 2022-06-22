package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
)

func UpdateRoomHandler(log configs.Logger, roomService chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		roomToUpdate := &Room{}

		if err := c.BodyParser(roomToUpdate); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		domainRoom := roomToDomainRoom(*roomToUpdate)
		usr := c.Context().UserValue("user")
		u := usr.(*user.User)

		if _, err := roomService.UpdateRoom(u.ID, domainRoom); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}
}
