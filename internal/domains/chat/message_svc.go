package chat

type MessageService interface {
	GetMessages(limit, offset, userID, chatID int64) ([]Message, error) // get certain messages in room
	CreateMessage(message Message) (*Message, error)
	UpdateMessage(message Message) (*Message, error)     // change content of message
	ReadMessage(messageID, userID int64) (bool, error)   // update with read status
	DeleteMessage(userID, messageID int64) (bool, error) // just update deleted_user list
	DeleteMyMessage(messageID int64) (bool, error)       // totally delete
}

type MessageServiceImpl struct {
	messageService MessageService
}

func (m *MessageServiceImpl) GetMessages(limit, offset, userID, chatID int64) ([]Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageServiceImpl) CreateMessage(message Message) (*Message, error) {
	//TODO implement me
	panic("implement me ")
}

func (m *MessageServiceImpl) UpdateMessage(message Message) (*Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageServiceImpl) ReadMessage(messageID, userID int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageServiceImpl) DeleteMessage(userID, messageID int64) (bool, error) {
	ok, err := m.messageService.DeleteMessage(userID, messageID)
	if err != nil {
		return false, err
	}
	return ok, err
}

func (m *MessageServiceImpl) DeleteMyMessage(messageID int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewMessageServiceImpl(messageService MessageService) *MessageServiceImpl {
	return &MessageServiceImpl{messageService: messageService}
}
