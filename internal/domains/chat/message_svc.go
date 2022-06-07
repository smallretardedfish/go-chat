package chat

import "github.com/smallretardedfish/go-chat/internal/repositories/message_repo"

type MessageService interface {
	GetMessages(limit, offset, userID, chatID int64) ([]Message, error) // get certain messages in room
	CreateMessage(message Message) (*Message, error)
	UpdateMessage(message Message) (*Message, error)     // change content of message
	ReadMessage(messageID, userID int64) (bool, error)   // update with read status
	DeleteMessage(userID, messageID int64) (bool, error) // just update deleted_user list
	DeleteMyMessage(messageID int64) (bool, error)       // totally delete
}

type MessageServiceImpl struct {
	messageRepo message_repo.MessageRepo
}

func (m *MessageServiceImpl) GetMessages(limit, offset, userID, roomID int64) ([]Message, error) {

	messages, err := m.messageRepo.GetMessages(&message_repo.MessageFilter{
		Search: nil,
		Limit:  &limit,
		Offset: &offset,
	}, userID, roomID)

	if err != nil {
		return nil, err
	}

	var result []Message
	for i := range messages {
		msg := repoMessageToDomainMessage(messages[i])
		result = append(result, msg)
	}

	return result, nil
}

func (m *MessageServiceImpl) CreateMessage(message Message) (*Message, error) {
	repoMsg := domainMessageToRepoMessage(message)
	msg, err := m.messageRepo.CreateMessage(repoMsg)
	if err != nil {
		return nil, err
	}
	res := repoMessageToDomainMessage(*msg)
	return &res, nil
}

func (m *MessageServiceImpl) UpdateMessage(message Message) (*Message, error) {
	repoMsg := domainMessageToRepoMessage(message)
	msg, err := m.messageRepo.UpdateMessage(repoMsg)
	if err != nil {
		return nil, err
	}
	res := repoMessageToDomainMessage(*msg)
	return &res, nil
}

func (m *MessageServiceImpl) ReadMessage(messageID, userID int64) (bool, error) { // ??
	//TODO implement me
	panic("implement me")
}

func (m *MessageServiceImpl) DeleteMessage(userID, messageID int64) (bool, error) {
	ok, err := m.messageRepo.DeleteMessageForUser(userID, messageID)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (m *MessageServiceImpl) DeleteMyMessage(messageID int64) (bool, error) {
	ok, err := m.messageRepo.DeleteMessage(messageID)
	if err != nil {
		return false, err
	}
	return ok, err
}

func NewMessageServiceImpl(repo message_repo.MessageRepo) *MessageServiceImpl {
	return &MessageServiceImpl{messageRepo: repo}
}
