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

type RoomUpdateData struct {
	Name *string `json:"name"`
	Type *int8   `json:"type"`
}

//TODO research how to implement this in more convenient way
// should it be in this layer of application???
func updateRoom(data RoomUpdateData, room chat.Room) *chat.Room { // for partial room update

	if data.Name != nil {
		room.Name = *data.Name
	}

	if data.Type != nil {
		room.Type = *data.Type
	}
	return &room
}

//TODO investigate whether are 2 queries is OK for update
func UpdateRoomHandler(log configs.Logger, roomService chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		roomID, err := strconv.Atoi(idStr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		log.Info("TRYING TO UPDATE room:", roomID)

		usr := c.Context().UserValue("user")
		u, ok := usr.(*user.User)
		if !ok {
			return fmt.Errorf("error: can't get user from token")
		}

		room, err := roomService.GetRoom(u.ID, int64(roomID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		if room == nil {
			return fmt.Errorf("user doesn't belong to this room or room doesn't exist")
		}

		data := &RoomUpdateData{}
		if err = c.BodyParser(data); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}

		updatedRoom := updateRoom(*data, *room)

		if _, err := roomService.UpdateRoom(u.ID, *updatedRoom); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}
}
