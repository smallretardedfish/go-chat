package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
)

// TODO make mapper from repo layer to service

func repoMessageToServiceMessage(message message_repo.Message) Message {
	return Message{
		ID:   message.ID,
		Text: message.Text,
		Owner: User{
			Name: message.Owner.Name,
		},
	}
}

func serviceMessageToRepoMessage(message Message) message_repo.Message { // TODO implement proper mapping
	return message_repo.Message{
		ID:      message.ID,
		Text:    message.Text,
		OwnerID: message.Owner.ID,
	}
}
