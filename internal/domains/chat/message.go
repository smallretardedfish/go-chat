package chat

import (
	"time"
)

type Message struct {
	ID        int64
	Text      string
	OwnerID   int64
	Owner     User
	RoomID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
