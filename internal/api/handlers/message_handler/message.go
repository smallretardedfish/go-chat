package message_handler

import (
	"github.com/smallretardedfish/go-chat/internal/api/handlers/common"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"time"
)

type Message struct {
	ID        int64       `json:"id"`
	Text      string      `json:"text"`
	Owner     common.User `json:"owner"`
	RoomID    int64       `json:"room_id"`
	IsRead    bool        `json:"is_read"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func domainMessageToMessage(message Message) chat.Message {
	return chat.Message{
		ID:        message.ID,
		Text:      message.Text,
		OwnerID:   message.Owner.ID,
		Owner:     common.UserToChatUser(message.Owner),
		RoomID:    message.RoomID,
		IsRead:    message.IsRead,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}

func chatMessageToMessage(message chat.Message) Message {
	return Message{
		ID:        message.ID,
		Text:      message.Text,
		Owner:     common.ChatUserToUser(message.Owner),
		RoomID:    message.RoomID,
		IsRead:    message.IsRead,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}
