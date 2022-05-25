package message_repo

import (
	"gorm.io/gorm"
)

type MessageRepo interface {
	CreateMessage(message Message) (*Message, error)
	MessagesByUser(userID int64) ([]*Message, error)
	AllMessagesByUserInRoom(userID int64, roomID int64) ([]*Message, error)
	AllMessagesInRoom(roomId int64) ([]*Message, error)
	UpdateMessage(userID, roomID int64, message Message) (*Message, error)
}

type MessageRepoPG struct {
	db *gorm.DB
}

func (m MessageRepoPG) CreateMessage(msg Message) (*Message, error) {
	res := m.db.Create(&msg)
	return &msg, res.Error
}

func (m MessageRepoPG) MessagesByUser(userID int64) ([]*Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepoPG) AllMessagesByUserInRoom(userID int64, roomID int64) ([]*Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepoPG) AllMessagesInRoom(roomId int64) ([]*Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepoPG) UpdateMessage(userID, roomID int64, message Message) (*Message, error) {
	//TODO implement me
	panic("implement me")
}

func NewMessageRepo(db *gorm.DB) *MessageRepoPG {
	return &MessageRepoPG{
		db: db,
	}
}
