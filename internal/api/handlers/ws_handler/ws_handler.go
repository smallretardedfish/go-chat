package ws_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/smallretardedfish/go-chat/internal/connector"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
)

func WsHandler(log logging.Logger, conn connector.Connector) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return websocket.New(func(wsConn *websocket.Conn) {
			usr := c.Context().UserValue("user").(*user.User)
			newWsConnection := connector.NewWsConnection(wsConn, connector.DomainUserToUser(*usr))
			conn.AddConnection(newWsConnection)
		})(c)
	}
}
