package room_handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/tools/slice"
	"net/http"
	"strconv"
)

//TODO reseach error handler
func GetRoomsHandler(log configs.Logger, service chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		limitStr := c.Query("limit", "100")
		offsetStr := c.Query("offset", "0")
		limit, err := strconv.Atoi(limitStr)
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		log.Info(c.Context().UserValue("user"))
		usr, ok := c.Context().UserValue("user").(*user.User)
		log.Info(usr)
		if !ok {
			return fmt.Errorf("error: can`t assert user from context to *user.User")
		}
		domainRooms, err := service.GetRooms(int64(limit), int64(offset), usr.ID)
		rooms := slice.Map(domainRooms, domainRoomToRoom)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return c.JSON(rooms)
	}
}
