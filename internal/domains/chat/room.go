package chat

import (
	"github.com/smallretardedfish/go-chat/internal/domains/user"
)

type Room struct {
	ID    int64
	Users []user.User
}
