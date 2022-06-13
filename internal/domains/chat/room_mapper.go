package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"github.com/smallretardedfish/go-chat/tools/slice"
)

// TODO refactor with slice.Map

func repoUserToUser(repoUser user_repo.User) User {
	return User{
		ID:   repoUser.ID,
		Name: repoUser.Name,
	}
}

func repoRoomToRoom(room room_repo.Room) Room {
	return Room{
		ID:      room.ID,
		Name:    room.Name,
		OwnerID: room.OwnerID,
		Owner: User{
			ID:   room.Owner.ID,
			Name: room.Owner.Name,
		},
		Type:      room.Type,
		Users:     slice.Map(room.Users, repoUserToUser),
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
}

func RoomToRepoRoom(room Room) room_repo.Room {
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
