package connector

import (
	"fmt"
	"github.com/smallretardedfish/go-chat/logging"
	"sync"
)

type Connector interface {
	AddConnection(conn Connection)
}

type ConnectorImpl struct {
	sync.RWMutex
	log             logging.Logger
	userConnections map[int64]Connection
}

func NewConnector(log logging.Logger) *ConnectorImpl {
	return &ConnectorImpl{
		log:             log,
		userConnections: make(map[int64]Connection),
	}
}

func (c *ConnectorImpl) AddConnection(conn Connection) {
	c.connect(conn)
}

func (c *ConnectorImpl) connect(conn Connection) {
	conn.Connect()
	c.userConnections[conn.GetUser().ID] = conn
	go c.listen(conn)
}

func (c *ConnectorImpl) listen(conn Connection) {
	for {
		select {
		case msg := <-conn.GetMessageChannel():
			fmt.Println(msg)
			conn.SendMessage(msg)
		}
	}
}
