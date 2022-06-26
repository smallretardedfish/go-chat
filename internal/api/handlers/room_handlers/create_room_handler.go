package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"net/http"
)

type roomCreationData struct {
	Room    Room    `json:"room"`
	Members []int64 `json:"members"`
}

func CreateRoomHandler(log logging.Logger, roomService chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data := &roomCreationData{}
		err := c.BodyParser(data)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		domainRoom := roomToDomainRoom(data.Room)
		ownr := c.Context().UserValue("user")
		owner := ownr.(*user.User)
		domainRoom.OwnerID = owner.ID
		if _, err := roomService.CreateRoom(domainRoom, data.Members); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}
}
