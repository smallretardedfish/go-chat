package message_repo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

//TODO fix Preloading of fields

type MessageRepo interface {
	CreateMessage(message Message) (*Message, error)
	GetMessage(messageID int64) (*Message, error)
	GetMessages(messageFilter *MessageFilter, userID, roomID int64) ([]Message, error)
	UpdateMessage(message Message) (*Message, error)
	DeleteMessage(messageID int64) (bool, error)
	DeleteMessageForUser(messageID, userID int64) (bool, error)
}

var _ MessageRepo = (*MessageRepoPG)(nil)

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

func (m *MessageRepoPG) GetMessage(messageID int64) (*Message, error) {
	message := Message{}
	err := m.db.Model(Message{}).Where("id = ?", messageID).First(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (m *MessageRepoPG) GetMessages(messageFilter *MessageFilter, userID, roomID int64) ([]Message, error) {
	var messages []Message
	res := m.db.Where("room_id = ? AND ? != ALL(deleted_users) ", roomID, userID).Find(&messages) //TODO implement logic of showing non-deleted messages
	if messageFilter != nil {
		if messageFilter.Search != nil {
			res = res.Where("text LIKE ?", fmt.Sprintf("%%%s%%", *messageFilter.Search))
		}
		if messageFilter.Offset != nil {
			res = res.Offset(int(*messageFilter.Offset))
		}
		if messageFilter.Limit != nil {
			res = res.Limit(int(*messageFilter.Limit))
		}
	}
	err := res.Error
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

func (m *MessageRepoPG) DeleteMessage(messageID int64) (bool, error) {
	err := m.db.Delete(Message{}, messageID).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *MessageRepoPG) DeleteMessageForUser(messageID, userID int64) (bool, error) {
	err := m.db.Exec("UPDATE messages SET deleted_users = array_append(deleted_users,?) WHERE message_id = ?", userID, messageID).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewMessageRepo(db *gorm.DB) *MessageRepoPG {
	return &MessageRepoPG{
		db: db,
	}
}
