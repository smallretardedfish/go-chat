package room_handlers

import (
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"time"
)

type Room struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	OwnerID   int64     `json:"owner_id"`
	Type      int8      `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func roomToDomainRoom(room Room) chat.Room {
	return chat.Room{
		ID:        room.ID,
		Name:      room.Name,
		OwnerID:   room.OwnerID,
		Type:      room.Type,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
}

func domainRoomToRoom(room chat.Room) Room {
	return Room{
		ID:        room.ID,
		Name:      room.Name,
		OwnerID:   room.OwnerID,
		Type:      room.Type,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
}
