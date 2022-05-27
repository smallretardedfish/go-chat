package message_repo

import (
	"gorm.io/gorm"
)

type MessageRepo interface {
	CreateMessage(message Message, roomID int64) (*Message, error)
	MessagesByUser(userID int64) ([]Message, error)
	AllMessagesByUserInRoom(userID int64, roomID int64) ([]Message, error)
	AllMessagesInRoom(roomId int64) ([]*Message, error)
	UpdateMessage(userID, roomID int64, message Message) (Message, error)
}

type MessageRepoPG struct {
	db *gorm.DB
}

func (m MessageRepoPG) CreateMessage(msg Message) (*Message, error) { // TODO: make it a  gorm transaction
	res := m.db.Create(&msg).Create(UserMessage{
		UserID: msg.OwnerID,
		Status: 1, // unread
	})

	return &msg, res.Error
}

func (m MessageRepoPG) MessagesByUser(userID int64) ([]Message, error) {
	var messages []Message
	res := m.db.Model(UserMessage{}).Where("user_id = ?", userID).Find(&messages)
	return messages, res.Error
}

func (m MessageRepoPG) AllMessagesByUserInRoom(userID int64, roomID int64) ([]Message, error) {
	var messages []Message
	res := m.db.Where("user_id = ? AND room_id = ?", userID, roomID).Find(&messages)
	return messages, res.Error
}

func (m MessageRepoPG) AllMessagesInRoom(roomId int64) ([]*Message, error) {
	var messages []*Message
	res := m.db.Where("room_id = ?", roomId).Find(&messages)
	return messages, res.Error
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
