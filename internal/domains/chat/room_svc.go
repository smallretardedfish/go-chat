package chat

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
	"github.com/smallretardedfish/go-chat/tools/slice"
)

type RoomService interface {
	GetRoom(userID, roomID int64) (*Room, error)
	GetRooms(limit, offset, userID int64) ([]Room, error)
	CreateRoom(room Room, userIDs []int64) (*Room, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) (bool, error)
	AddUserToRoom(userID, roomID int64) (bool, error)
	DeleteUserFromRoom(userID, roomID int64) (bool, error)
}

type RoomServiceImpl struct {
	roomRepo room_repo.RoomRepo
}

func (r *RoomServiceImpl) GetRoom(userID, roomID int64) (*Room, error) {
	room, err := r.roomRepo.GetRoom(userID, roomID)
	if err != nil || room == nil {
		return nil, err
	}

	res := repoRoomToRoom(*room)
	return &res, nil
}

func (r *RoomServiceImpl) GetRooms(limit, offset, userID int64) ([]Room, error) {
	rooms, err := r.roomRepo.GetRooms(userID, &room_repo.RoomFilter{
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {
		return nil, err
	}
	return slice.Map(rooms, repoRoomToRoom), nil
}

func (r *RoomServiceImpl) CreateRoom(room Room, userIDs []int64) (*Room, error) {
	repoRoom := RoomToRepoRoom(room)
	createdRoom, err := r.roomRepo.CreateRoom(repoRoom)
	if err != nil {
		return nil, err
	}
	members := []room_repo.RoomUser{{RoomID: createdRoom.ID, UserID: createdRoom.OwnerID}} // first member of room is owner

	for _, userID := range userIDs {
		member := room_repo.RoomUser{
			RoomID: createdRoom.ID,
			UserID: userID,
		}
		members = append(members, member)
	}

	if _, err := r.roomRepo.CreateRoomUsers(members); err != nil { // inserting all members together
		return nil, err
	}
	res := repoRoomToRoom(*createdRoom)
	return &res, nil
}

func (r *RoomServiceImpl) UpdateRoom(userID int64, room Room) (*Room, error) {
	repoRoom := RoomToRepoRoom(room)
	updatedRoom, err := r.roomRepo.UpdateRoom(userID, repoRoom)
	if err != nil {
		return nil, err
	}
	res := repoRoomToRoom(*updatedRoom)
	return &res, nil
}

func (r *RoomServiceImpl) DeleteRoom(userID, roomID int64) (bool, error) {
	ok, err := r.roomRepo.DeleteRoom(roomID, userID)
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
