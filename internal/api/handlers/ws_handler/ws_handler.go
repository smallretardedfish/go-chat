package ws_handler

import (
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"github.com/smallretardedfish/go-chat/internal/connector"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(ctx *fasthttp.RequestCtx) bool { return true },
}

func WsHandler(log logging.Logger, conn connector.Connector) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		usr := ctx.UserValue("user").(*user.User)
		if err := upgrader.Upgrade(ctx, func(wsConn *websocket.Conn) {
			newWsConnection := connector.NewWsConnection(wsConn, connector.DomainUserToUser(*usr))
			log.Debug("new wsConn ", newWsConnection)
			conn.AddConnection(newWsConnection)
		}); err != nil {
			log.Error(err)
			return err
		}
		return nil
	}
}
