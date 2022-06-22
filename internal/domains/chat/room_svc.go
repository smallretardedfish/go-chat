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
	AddUsersToRoom(userIDs []int64, roomID int64) (bool, error)
	DeleteUsersFromRoom(initiatorID, roomID int64, userID []int64) (bool, error) // owner deletes batch of users
	DeleteCurrentUser(userID, roomID int64) (bool, error)                        // user deletes themself from room
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

	if err := r.roomRepo.CreateRoomUsers(members); err != nil { // inserting all members together
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

func (r *RoomServiceImpl) AddUsersToRoom(userIDs []int64, roomID int64) (bool, error) {

	var RoomUsers []room_repo.RoomUser
	for _, id := range userIDs {
		RoomUsers = append(RoomUsers, room_repo.RoomUser{
			RoomID: roomID,
			UserID: id,
		})
	}
	if err := r.roomRepo.CreateRoomUsers(RoomUsers); err != nil {
		return false, err
	}
	return true, nil
}

func (r *RoomServiceImpl) DeleteCurrentUser(userID, roomID int64) (bool, error) {
	ok, err := r.roomRepo.DeleteCurrentRoomUser(userID, roomID)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (r *RoomServiceImpl) DeleteUsersFromRoom(initiatorID, roomID int64, toRemove []int64) (bool, error) {
	ok, err := r.roomRepo.DeleteRoomUsers(initiatorID, roomID, toRemove)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func NewRoomServiceImpl(roomRepo room_repo.RoomRepo) *RoomServiceImpl {
	return &RoomServiceImpl{roomRepo: roomRepo}
}
