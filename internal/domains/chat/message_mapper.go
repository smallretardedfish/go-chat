package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
)

func repoMessageToMessage(message message_repo.Message) Message {
	return Message{
		ID:      message.ID,
		Text:    message.Text,
		OwnerID: message.OwnerID,
		Owner: User{
			ID:   message.OwnerID,
			Name: message.Owner.Name,
		},
		RoomID:    message.RoomID,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}

func messageToRepoMessage(message Message) message_repo.Message {
	return message_repo.Message{
		ID:        message.ID,
		Text:      message.Text,
		OwnerID:   message.OwnerID,
		RoomID:    message.RoomID,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}
