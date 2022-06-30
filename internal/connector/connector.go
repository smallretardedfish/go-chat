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
}

type ConnectorImpl struct {
	sync.RWMutex
	log             logging.Logger
	userConnections map[int64]Connection
	rooms           map[int64][]int64 // room and active users inside it
	messageSvc      chat.MessageService
}

func NewConnector(log logging.Logger, messageSvc chat.MessageService) *ConnectorImpl {
	return &ConnectorImpl{
		log:             log,
		userConnections: make(map[int64]Connection),
		rooms:           make(map[int64][]int64),
		messageSvc:      messageSvc,
	}
}

func (c *ConnectorImpl) AddConnection(conn Connection) {
	c.connect(conn)
}

func (c *ConnectorImpl) connect(conn Connection) {
	go conn.Connect()

	c.Lock()
	c.userConnections[conn.GetUser().ID] = conn
	c.Unlock()

	c.listen(conn)
}

func (c *ConnectorImpl) deleteUserFromRoom(roomID, userID int64) {
	if _, ok := c.rooms[roomID]; !ok {
		return
	}

	index := slices.Index(c.rooms[roomID], userID)
	if index == -1 {
		return
	}

	c.rooms[roomID] = slices.Delete(c.rooms[roomID], index, index+1)
}

func (c *ConnectorImpl) addUserToRoom(roomID, userID int64) {
	c.Lock()
	c.rooms[roomID] = append(c.rooms[roomID], userID) // works fine
	c.Unlock()
}

func (c *ConnectorImpl) deleteUserConnection(userID int64) {
	c.Lock()
	delete(c.userConnections, userID)

	c.Unlock()
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
	//
	err := c.SendMessageToRoom(conn.GetRoom(), conn.GetUser().Username+": "+data.Text)
	if err != nil {
		c.log.Error(err)
		return
	}
}

func (c *ConnectorImpl) onDisconnect(conn Connection) {
	c.Lock()
	c.deleteUserConnection(conn.GetUser().ID)
	c.deleteUserFromRoom(conn.GetRoom(), conn.GetUser().ID)
	c.Unlock()
	err := conn.Close()
	c.log.Debug(conn.GetUser().Username, "left from room:", conn.GetRoom())
	if err != nil {
		c.log.Error(err)
		return
	}
}

func (c *ConnectorImpl) SendMessageToRoom(roomID int64, msg any) error { //kinda broadcast
	for _, userID := range c.rooms[roomID] {
		if err := c.userConnections[userID].SendMessage(map[string]interface{}{"message": msg}); err != nil {
			return err
		}
	}
	return nil
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
