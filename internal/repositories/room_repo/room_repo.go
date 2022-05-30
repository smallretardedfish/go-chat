package room_repo

import (
	"gorm.io/gorm"
)

type RoomRepo interface {
	GetRoom(userID, roomID int64) (*Room, error)
	GetRooms(userID int64) ([]Room, error)
	CreateRoom(room Room) (*Room, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) error
}

type RoomRepoPG struct {
	db *gorm.DB
}

func (r RoomRepoPG) GetRoom(userID, roomID int64) (*Room, error) {
	roomUser := RoomUser{}
	res := r.db.Model(RoomUser{}).Where("user_id = ? AND room_id = ?", userID, roomID).First(&roomUser)
	if res.Error != nil {
		return nil, res.Error
	}
	room := Room{}
	res = r.db.Model(Room{}).Preload("RoomUsers").Where("id = ?", roomUser.RoomID).First(&room)
	return &room, res.Error
}

func (r RoomRepoPG) GetRooms(userID int64) ([]Room, error) { // TODO fix gorm belongs-to population of RoomUsers slice
	var rooms []Room
	res := r.db.Table("rooms r").
		Joins(`JOIN room_users ru ON r.id = ru.room_id`).
		Where("ru.user_id = ?", userID).Preload("Owner").
		Preload("RoomUsers").Scan(&rooms)

	return rooms, res.Error
}

func (r RoomRepoPG) CreateRoom(room Room) (*Room, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error { // for consistency (room should be created before room_user))
		res := r.db.Create(&room)
		if res.Error != nil {
			return res.Error
		}
		roomUser := RoomUser{
			RoomID: room.ID,
			UserID: room.OwnerID,
			Status: RoomUserCreated,
		}
		res = r.db.Create(&roomUser)

		return res.Error
	})
	return &room, err
}

func (r RoomRepoPG) UpdateRoom(userID int64, room Room) (*Room, error) {
	roomUser := RoomUser{}
	res := r.db.Model(RoomUser{}).Where("user_id = ? AND room_id = ?", userID, room.ID).First(&roomUser)
	if res.Error != nil {
		return nil, res.Error
	}
	res = r.db.Model(Room{}).Where("id = ?", roomUser.RoomID).Save(&room)

	return &room, res.Error
}

func (r RoomRepoPG) DeleteRoom(userID, roomID int64) error {
	res := r.db.Model(RoomUser{}).Where("user_id = ? AND room_id = ?", userID, roomID).Update("status", RoomUserDeleted)
	return res.Error
}

func NewRoomRepo(db *gorm.DB) *RoomRepoPG {
	return &RoomRepoPG{
		db: db,
	}
}
