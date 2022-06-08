package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
)

//TODO investigate how to map user list of room in one direction and vice versa

func repoRoomToDomainRoom(room room_repo.Room) Room {
	var users []User

	for _, user := range room.Users {
		users = append(users, User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return Room{
		ID:      room.ID,
		Name:    room.Name,
		OwnerID: room.OwnerID,
		Owner: User{
			ID:   room.Owner.ID,
			Name: room.Owner.Name,
		},
		Type:      room.Type,
		Users:     users,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
}

func domainRoomToRepoRoom(room Room) room_repo.Room {
	return room_repo.Room{
		ID:        room.ID,
		Name:      room.Name,
		OwnerID:   room.OwnerID,
		Type:      room.Type,
		Users:     nil, // ??
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
}
