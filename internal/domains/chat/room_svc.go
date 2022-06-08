package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
	"log"
)

type RoomService interface {
	GetRoom(userID, roomID int64) (*Room, error)
	GetRooms(limit, offset, userID int64) ([]Room, error)
	GetUsers(roomID, userID int64) ([]User, error)
	CreateRoom(room Room, userIDs []int64) (*Room, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) (bool, error) // TODO research whether is this total deletion of room??
	AddUserToRoom(userID, roomID int64) (bool, error)
	DeleteUserFromRoom(userID, roomID int64) (bool, error)
}

type RoomServiceImpl struct {
	roomRepo room_repo.RoomRepo
}

func (r *RoomServiceImpl) GetRoom(userID, roomID int64) (*Room, error) {
	room, err := r.roomRepo.GetRoom(userID, roomID)
	if err != nil {
		return nil, err
	}
	log.Println("REPO ROOM OWNER NAME IS:", room.Owner.Name)
	log.Println("REPO ROOM USERS ARE:", room.Users)

	res := repoRoomToDomainRoom(*room)
	return &res, nil
}

func (r *RoomServiceImpl) GetRooms(limit, offset, userID int64) ([]Room, error) {
	rooms, err := r.roomRepo.GetRooms(userID)
	if err != nil {
		return nil, err
	}
	var res []Room
	for i := range rooms {
		room := repoRoomToDomainRoom(rooms[i])
		res = append(res, room)
	}
	return res, nil
}

func (r *RoomServiceImpl) GetUsers(roomID, userID int64) ([]User, error) {
	room, err := r.GetRoom(userID, roomID)
	if err != nil {
		return nil, err
	}
	return room.Users, nil
}

func (r *RoomServiceImpl) CreateRoom(room Room, userIDs []int64) (*Room, error) {
	repoRoom := domainRoomToRepoRoom(room)
	createdRoom, err := r.roomRepo.CreateRoom(repoRoom)
	if err != nil {
		return nil, err
	}
	for _, userID := range userIDs {
		_, err := r.AddUserToRoom(userID, createdRoom.ID)
		if err != nil {
			return nil, err
		}
	}
	res := repoRoomToDomainRoom(*createdRoom)
	return &res, nil
}

func (r *RoomServiceImpl) UpdateRoom(userID int64, room Room) (*Room, error) {
	repoRoom := domainRoomToRepoRoom(room)
	updatedRoom, err := r.roomRepo.UpdateRoom(userID, repoRoom)
	if err != nil {
		return nil, err
	}
	res := repoRoomToDomainRoom(*updatedRoom)
	return &res, nil
}

func (r *RoomServiceImpl) DeleteRoom(userID, roomID int64) (bool, error) {
	ok, err := r.roomRepo.DeleteRoomUser(roomID, userID)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (r *RoomServiceImpl) AddUserToRoom(userID, roomID int64) (bool, error) {
	_, err := r.roomRepo.CreateRoomUser(room_repo.RoomUser{ // should I create roomUser like this?
		RoomID: roomID,
		UserID: userID,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RoomServiceImpl) DeleteUserFromRoom(userID, roomID int64) (bool, error) {
	ok, err := r.roomRepo.DeleteRoomUser(roomID, userID)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func NewRoomServiceImpl(roomRepo room_repo.RoomRepo) *RoomServiceImpl {
	return &RoomServiceImpl{roomRepo: roomRepo}
}
