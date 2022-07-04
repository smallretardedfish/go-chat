package message_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
	"github.com/smallretardedfish/go-chat/pkg/slice"
	"net/http"
	"strconv"
)

func GetMessagesHandler(log logging.Logger, service chat.MessageService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		limitStr := c.Query("limit", "10")
		offsetStr := c.Query("offset", "0")
		roomIdStr := c.Params("room_id")
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		offset, err := strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		roomID, err := strconv.ParseInt(roomIdStr, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
		usr := c.Context().UserValue("user").(*user.User)
		domainMessages, err := service.GetMessages(usr.ID, roomID, &limit, &offset)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		messages := slice.Map(domainMessages, chatMessageToMessage)

		return c.JSON(messages)
	}
}
