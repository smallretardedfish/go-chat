package connector

import (
	"encoding/json"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/logging"
	"golang.org/x/exp/slices"
	"sync"
	"time"
)

type Connector interface {
	AddConnection(conn Connection)
	SendMessageToRoom(roomID int64, msg any)
	SendMessage(userID int64, msg any)
}

type ConnectorImpl struct {
	sync.RWMutex
	log             logging.Logger
	userConnections map[int64][]Connection
	rooms           map[int64]UserSet // room and active users inside it
	messageSvc      chat.MessageService
}

type UserSet map[int64]struct{}

func (us UserSet) Add(userID int64) {
	us[userID] = struct{}{}
}

func (us UserSet) Remove(userID int64) {
	delete(us, userID)
}

func NewConnector(log logging.Logger, messageSvc chat.MessageService) *ConnectorImpl {
	return &ConnectorImpl{
		log:             log,
		userConnections: make(map[int64][]Connection),
		rooms:           make(map[int64]UserSet),
		messageSvc:      messageSvc,
	}
}

func (c *ConnectorImpl) AddConnection(conn Connection) {
	c.connect(conn)
}

func (c *ConnectorImpl) connect(conn Connection) {
	go conn.Connect()

	c.addUserConnection(conn)

	c.listen(conn)
}

func (c *ConnectorImpl) addUserToRoom(roomID, userID int64) {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.rooms[roomID]; !ok {
		c.rooms[roomID] = make(UserSet)
	}
	c.rooms[roomID].Add(userID)

}

func (c *ConnectorImpl) deleteUserFromRoom(roomID, userID int64) {
	c.Lock()
	defer c.Unlock()
	c.rooms[roomID].Remove(userID)
}

func (c *ConnectorImpl) addUserConnection(conn Connection) {
	c.Lock()
	defer c.Unlock()
	c.userConnections[conn.GetUser().ID] = append(c.userConnections[conn.GetUser().ID], conn)
}

func (c *ConnectorImpl) deleteUserConnection(conn Connection) {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.userConnections[conn.GetUser().ID]; !ok {
		return
	}

	index := slices.IndexFunc(c.userConnections[conn.GetUser().ID], func(c Connection) bool {
		return conn == c
	})
	if index == -1 {
		return
	}

	c.userConnections[conn.GetUser().ID] = slices.Delete(c.userConnections[conn.GetUser().ID], index, index+1)
}

func (c *ConnectorImpl) onMessage(conn Connection, data []byte) {
	eventMessage := &EventMessage{}

	err := json.Unmarshal(data, eventMessage)
	if err != nil {
		c.log.Error(err)
		return
	}
	switch eventMessage.Type {
	case ConnectToRoomEventMessageType:
		c.onConnectToRoom(conn, eventMessage.Data.ConnectToRoomData)
	case NewMessageEventMessageType:
		c.onNewMessage(conn, eventMessage.Data.NewMessageData)
	}

}

func (c *ConnectorImpl) onConnectToRoom(conn Connection, data *ConnectToRoomData) {
	c.addUserToRoom(data.RoomID, conn.GetUser().ID)
	conn.SetRoom(data.RoomID)
}

func (c *ConnectorImpl) onNewMessage(conn Connection, data *NewMessageData) {
	if _, err := c.messageSvc.CreateMessage(chat.Message{
		Text:      data.Text,
		OwnerID:   conn.GetUser().ID,
		RoomID:    conn.GetRoom(),
		IsRead:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		c.log.Error(err)
		return
	}
	msg := EventMessage{
		Type: NewMessageEventMessageType,
		Data: Data{
			NewMessageData: data,
		},
	}
	c.SendMessageToRoom(conn.GetRoom(), msg)

}

func (c *ConnectorImpl) onDisconnect(conn Connection) {
	c.deleteUserConnection(conn)
	c.deleteUserFromRoom(conn.GetRoom(), conn.GetUser().ID)

	err := conn.Close()
	c.log.Debug(conn.GetUser().Username, "left from room:", conn.GetRoom())
	if err != nil {
		c.log.Error(err)
		return
	}
}

func (c *ConnectorImpl) SendMessage(userID int64, msg any) {
	conns := c.userConnections[userID]
	for _, conn := range conns {
		if err := conn.SendMessage(msg); err != nil {
			c.log.Error(err)
		}
	}
}

func (c *ConnectorImpl) SendMessageToRoom(roomID int64, msg any) { //kinda broadcast
	for userID := range c.rooms[roomID] {
		c.SendMessage(userID, msg)
	}
}

func (c *ConnectorImpl) listen(conn Connection) {
	for {
		select {
		case msg := <-conn.GetMessageChannel():
			c.onMessage(conn, msg)
		case <-conn.GetDisconnectClientCh():
			c.onDisconnect(conn)
		}
	}
}
