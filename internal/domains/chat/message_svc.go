package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
	"github.com/smallretardedfish/go-chat/tools/slice"
)

type MessageService interface {
	GetMessages(limit, offset, userID, chatID int64) ([]Message, error) // get certain messages in room
	CreateMessage(message Message) (*Message, error)
	UpdateMessage(message Message) (*Message, error)     // change content of message
	DeleteMessage(userID, messageID int64) (bool, error) // totally delete
	DeleteMyMessage(messageID int64) (bool, error)       // just update deleted_user list
}

type MessageServiceImpl struct {
	messageRepo message_repo.MessageRepo
}

func (m *MessageServiceImpl) GetMessages(limit, offset *int64, userID, roomID int64) ([]Message, error) {

	messages, err := m.messageRepo.GetMessages(&message_repo.MessageFilter{
		Limit:  limit,
		Offset: offset,
	}, userID, roomID)

	if err != nil {
		return nil, err
	}
	return slice.Map(messages, repoMessageToMessage), nil
}

func (m *MessageServiceImpl) CreateMessage(message Message) (*Message, error) {
	repoMsg := messageToRepoMessage(message)
	msg, err := m.messageRepo.CreateMessage(repoMsg)
	if err != nil {
		return nil, err
	}
	res := repoMessageToMessage(*msg)
	return &res, nil
}

func (m *MessageServiceImpl) UpdateMessage(message Message) (*Message, error) { //TODO fix deleted_users loss after mapping
	repoMsg := messageToRepoMessage(message)
	msg, err := m.messageRepo.UpdateMessage(repoMsg)
	if err != nil {
		return nil, err
	}
	res := repoMessageToMessage(*msg)
	return &res, nil
}

func (m *MessageServiceImpl) DeleteMessageForUser(userID, messageID int64) (bool, error) {
	ok, err := m.messageRepo.DeleteMessageForUser(userID, messageID)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (m *MessageServiceImpl) DeleteMessage(messageID int64) (bool, error) {
	ok, err := m.messageRepo.DeleteMessage(messageID)
	if err != nil {
		return false, err
	}
	return ok, err
}

func NewMessageServiceImpl(repo message_repo.MessageRepo) *MessageServiceImpl {
	return &MessageServiceImpl{messageRepo: repo}
}
