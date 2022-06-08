package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
)

// TODO make mapper from repo layer to service

func repoMessageToDomainMessage(message message_repo.Message) Message {
	return Message{
		ID:      message.ID,
		Text:    message.Text,
		OwnerID: message.OwnerID,
		Owner: User{
			ID:   message.OwnerID,
			Name: message.Owner.Name,
		},
		RoomID:       message.RoomID,
		DeletedUsers: message.DeletedUsers,
		CreatedAt:    message.CreatedAt,
		UpdatedAt:    message.UpdatedAt,
	}
}

func domainMessageToRepoMessage(message Message) message_repo.Message { // TODO implement proper mapping
	return message_repo.Message{
		ID:           message.ID,
		Text:         message.Text,
		OwnerID:      message.OwnerID,
		RoomID:       message.RoomID,
		DeletedUsers: message.DeletedUsers,
		CreatedAt:    message.CreatedAt,
		UpdatedAt:    message.UpdatedAt,
	}
}
