package chat

import (
	"time"
)

type Room struct {
	ID        int64
	Name      string
	OwnerID   int64
	Owner     User
	Type      int8
	Users     []User
	CreatedAt time.Time
	UpdatedAt time.Time
}
