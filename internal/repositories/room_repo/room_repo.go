package room_repo

import "gorm.io/gorm"

type RoomRepo interface {
	GetRoom(userID, roomID int64) (*Room, error)
	GetRooms(userID int64) ([]Room, error)
	CreateRoom(room Room) (*Room, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) error
}

type RoomRepoImpl struct {
	db *gorm.DB
}

func (r RoomRepoImpl) GetRoom(userID, roomID int64) (*Room, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoomRepoImpl) GetRooms(userID int64) ([]Room, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoomRepoImpl) CreateRoom(room Room) (*Room, error) {
	res := r.db.Create(&room)
	return &room, res.Error
}

func (r RoomRepoImpl) UpdateRoom(userID int64, room Room) (*Room, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoomRepoImpl) DeleteRoom(userID, roomID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewRoomRepo(db *gorm.DB) *RoomRepoImpl {
	return &RoomRepoImpl{
		db: db,
	}
}
