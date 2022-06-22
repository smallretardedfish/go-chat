package connector

import (
	"github.com/gofiber/websocket/v2"
)

type Connection interface {
	Connect()
	GetMessageChannel() chan []byte
	SendMessage(data any) error
	Close() error
	GetUser() *User
	GetRoom() int64
	SetRoom(id int64)
}

type WsConnection struct {
	conn          *websocket.Conn
	user          *User
	currentRoomId int64
	isConnected   bool
	messageCh     chan []byte
	closeCh       chan struct{}
}

func NewWsConnection(conn *websocket.Conn, user *User) *WsConnection {
	return &WsConnection{
		conn:      conn,
		user:      user,
		messageCh: make(chan []byte),
		closeCh:   make(chan struct{}),
	}
}

func (c *WsConnection) read() {
	for {
		select {
		case <-c.closeCh:
			c.isConnected = false
			return
		default:
			_, messageData, err := c.conn.ReadMessage()
			if err != nil {
				c.isConnected = false
				return
			}
			c.messageCh <- messageData
		}
	}

}
func (c *WsConnection) Connect() {
	go c.read()
}

func (c *WsConnection) GetMessageChannel() chan []byte {
	return c.messageCh
}

func (c *WsConnection) SendMessage(data any) error {
	return c.conn.WriteJSON(data)
}

func (c *WsConnection) Close() error {
	c.closeCh <- struct{}{}
	return c.conn.Close()
}

func (c *WsConnection) GetUser() *User {
	return c.user
}

func (c *WsConnection) GetRoom() int64 {
	return c.currentRoomId
}

func (c *WsConnection) SetRoom(id int64) {
	c.currentRoomId = id
}
