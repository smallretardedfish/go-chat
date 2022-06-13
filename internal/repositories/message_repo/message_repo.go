package message_repo

import (
	"fmt"
	"gorm.io/gorm"
)

type MessageRepo interface {
	CreateMessage(message Message) (*Message, error)
	GetMessage(messageID int64) (*Message, error)
	GetMessages(messageFilter *MessageFilter, userID, roomID int64) ([]Message, error)
	UpdateMessage(message Message) (*Message, error)
	DeleteMessage(messageID int64) (bool, error)
	DeleteMessageForUser(messageID, userID int64) (bool, error)
}

type MessageRepoPG struct {
	db *gorm.DB
}

func (m *MessageRepoPG) CreateMessage(msg Message) (*Message, error) {
	err := m.db.Model(Message{}).Create(&msg).Error
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (m *MessageRepoPG) GetMessage(messageID int64) (*Message, error) {
	message := Message{}
	err := m.db.Model(Message{}).Preload("Owner").Where("id = ?", messageID).First(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

//TODO check where clause
func (m *MessageRepoPG) GetMessages(messageFilter *MessageFilter, userID, roomID int64) ([]Message, error) {
	var messages []Message
	res := m.db.Preload("Owner").Where("room_id = ? AND ? != ALL(deleted_users) ", roomID, userID)
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
	err := res.Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (m *MessageRepoPG) UpdateMessage(message Message) (*Message, error) {
	err := m.db.Model(Message{}).Omit("deleted_users").Where("id = ?", message.ID).Save(message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (m *MessageRepoPG) DeleteMessage(messageID int64) (bool, error) {
	res := m.db.Delete(Message{}, messageID)
	err := res.Error
	if err != nil {
		return false, err
	}
	if res.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}

func (m *MessageRepoPG) DeleteMessageForUser(messageID, userID int64) (bool, error) {
	res := m.db.Exec("UPDATE messages SET deleted_users = array_append(deleted_users,?) WHERE id = ?", userID, messageID)
	err := res.Error
	if err != nil {
		return false, err
	}
	if res.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}

func NewMessageRepo(db *gorm.DB) *MessageRepoPG {
	return &MessageRepoPG{
		db: db,
	}
}
