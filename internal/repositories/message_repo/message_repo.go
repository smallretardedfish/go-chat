package message_repo

import (
	"errors"
	"gorm.io/gorm"
)

//TODO fix Preloading of fields

type MessageRepo interface {
	CreateMessage(message Message) (*Message, error)
	GetMessages(messageID int64) ([]Message, error)
	UpdateMessage(roomID int64, message Message) (*Message, error)
	DeleteMessage(messageID int64) (bool, error)
}

type MessageRepoPG struct {
	db *gorm.DB
}

func (m *MessageRepoPG) CreateMessage(msg Message) (*Message, error) { // TODO develop logic of creating message (store  people who deleted this message)
	err := m.db.Model(Message{}).Create(&msg).Error
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (m *MessageRepoPG) GetMessages(messageFilter *MessageFilter, userID, roomID int64) ([]Message, error) {
	var messages []Message
	err := m.db.Where("room_id = ? AND ? != ALL(deleted_users) ", roomID, userID).Find(&messages).Error //TODO implement logic of showing non-deleted messages
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return messages, err
}

func (m *MessageRepoPG) UpdateMessage(message Message) (*Message, error) {
	err := m.db.Model(Message{}).Where("id = ?", message.ID).Save(message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (m *MessageRepoPG) DeleteMessage(messageID int64) error {
	err := m.db.Delete(Message{}, messageID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func NewMessageRepo(db *gorm.DB) *MessageRepoPG {
	return &MessageRepoPG{
		db: db,
	}
}
