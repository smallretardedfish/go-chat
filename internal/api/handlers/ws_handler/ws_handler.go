package ws_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/connector"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
)

func WsHandler(log configs.Logger, conn connector.Connector) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return websocket.New(func(wsConn *websocket.Conn) {
			usr := c.Context().UserValue("user").(*user.User)
			newWsConnection := connector.NewWsConnection(wsConn, connector.DomainUserToUser(*usr))
			conn.AddConnection(newWsConnection)
		})(c)
	}
}
