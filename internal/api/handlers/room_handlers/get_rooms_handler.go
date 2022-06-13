package room_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/tools/slice"
	"net/http"
	"strconv"
)

//TODO reseach error handler
func GetRoomsHandler(log configs.Logger, service chat.RoomService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		limitStr := c.Query("limit", "100")
		offsetStr := c.Query("offset", "0")
		//	userIdStr := c.Query("id")
		limit, err := strconv.Atoi(limitStr)
		offset, err := strconv.Atoi(offsetStr)
		//userID, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		domainRooms, err := service.GetRooms(int64(limit), int64(offset), 1)
		rooms := slice.Map(domainRooms, domainRoomToRoom)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return c.JSON(rooms)
	}
}
