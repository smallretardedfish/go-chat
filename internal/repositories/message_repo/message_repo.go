package message_repo

import (
	"gorm.io/gorm"
)

type MessageRepo interface {
	CreateMessage(message Message) (*Message, error)
	GetMessage(messageID int64) (*Message, error)
	UpdateMessage(roomID int64, message Message) (*Message, error)
	DeleteMessageForUser(messageID int64) error
	DeleteMessageForAll(messageID int64) error
	MessagesByUser(userID int64) ([]Message, error)
	AllMessagesInRoom(roomId int64) ([]Message, error) // TODO: implement pagination
}

type MessageRepoPG struct {
	db *gorm.DB
}

func (m MessageRepoPG) GetMessage(messageID int64) (*Message, error) {
	msg := Message{}
	res := m.db.Model(Message{}).Preload("Owner").First(&msg, messageID)
	return &msg, res.Error
}

func (m MessageRepoPG) CreateMessage(msg Message) (*Message, error) {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		res := m.db.Model(Message{}).Create(&msg)
		if res.Error != nil {
			return res.Error
		}
		res = m.db.Model(UserMessage{}).Create(UserMessage{
			MessageID: msg.ID,
			UserID:    msg.OwnerID,
			Status:    UserMessageUnread, // unread
		})
		return res.Error
	})
	return &msg, err
}

func (m MessageRepoPG) MessagesByUser(userID int64) ([]Message, error) {
	var messages []Message
	res := m.db.Model(UserMessage{}).Where("owner_id = ?", userID).Find(&messages)
	return messages, res.Error
}

func (m MessageRepoPG) AllMessagesInRoom(roomID int64) ([]Message, error) { //userID is the initiator
	var messages []Message
	res := m.db.Where("room_id = ?", roomID).Find(&messages)
	return messages, res.Error
}

func (m MessageRepoPG) UpdateMessage(message Message) (*Message, error) {
	res := m.db.Model(Message{}).Where("id = ? AND room_id = ?", message.ID, message.RoomID).Updates(message)
	return &message, res.Error
}

func (m MessageRepoPG) DeleteMessageForAll(messageID int64) error {
	res := m.db.Where("id = ?", messageID).Delete(Message{})
	return res.Error
}

func (m MessageRepoPG) DeleteMessageForUser(messageID int64) error {
	res := m.db.Where("message_id = ?", messageID).Delete(UserMessage{})
	return res.Error
}

func NewMessageRepo(db *gorm.DB) *MessageRepoPG {
	return &MessageRepoPG{
		db: db,
	}
}
