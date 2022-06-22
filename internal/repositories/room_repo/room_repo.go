package room_repo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type RoomRepo interface {
	GetRoom(userID, roomID int64) (*Room, error)
	GetRooms(userID int64, roomFilter *RoomFilter) ([]Room, error)
	CreateRoom(room Room) (*Room, error)
	CreateRoomUsers(roomUser []RoomUser) error
	DeleteRoomUsers(initiatorID, roomID int64, toRemove []int64) (bool, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) (bool, error)
	DeleteCurrentRoomUser(roomID, userID int64) (bool, error)
}

type RoomRepoPG struct {
	db *gorm.DB
}

func (r *RoomRepoPG) GetRoom(userID, roomID int64) (*Room, error) {
	room := Room{}
	err := r.db.Preload("Owner").Preload("Users").
		Raw("SELECT * FROM rooms AS r JOIN room_users AS ru ON r.id = ru.room_id "+
			"WHERE ru.user_id = ? AND ru.room_id = ?", userID, roomID).
		First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepoPG) GetRooms(userID int64, roomFilter *RoomFilter) ([]Room, error) {
	var rooms []Room
	res := r.db.Preload("Owner").Preload("Users").
		Raw("SELECT * FROM rooms AS r JOIN room_users AS ru ON r.id = ru.room_id WHERE ru.user_id = ?", userID)

	if roomFilter != nil {
		if roomFilter.Search != nil {
			res = res.Where("text LIKE ?", fmt.Sprintf("%%%s%%", *roomFilter.Search))
		}
		if roomFilter.Offset != nil {
			res = res.Offset(int(*roomFilter.Offset))
		}
		if roomFilter.Limit != nil {
			res = res.Limit(int(*roomFilter.Limit))
		}
	}
	err := res.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

//non-native SQL is used to make batch insert
func (r *RoomRepoPG) CreateRoomUsers(roomUsers []RoomUser) error {
	err := r.db.Create(&roomUsers).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepoPG) DeleteCurrentRoomUser(roomID, userID int64) (bool, error) {
	res := r.db.Raw("DELETE FROM room_users AS ru WHERE ru.room = ? AND ru.user = ?",
		roomID, userID)
	err := res.Error
	if err != nil {
		return false, err
	}
	if res.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}

func (r *RoomRepoPG) DeleteRoomUsers(roomID, initiatorID int64, toRemove []int64) (bool, error) {
	res := r.db.Raw("DELETE FROM room_users AS ru INNER JOIN rooms AS r "+
		"ON ru.room_id = r.id  WHERE roomID = ? AND r.owner_id = ? AND ru.user_id IN ?",
		roomID, initiatorID, toRemove)
	err := res.Error
	if err != nil {
		return false, err
	}
	if res.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}

func (r *RoomRepoPG) CreateRoom(room Room) (*Room, error) {
	err := r.db.
		Raw("INSERT INTO rooms(owner_id,name) VALUES(?,?) RETURNING *", room.OwnerID, room.Name).
		Scan(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, err
}

func (r *RoomRepoPG) UpdateRoom(userID int64, room Room) (*Room, error) {
	//log.Println("UPDATING ROOM WITH ID", room.ID)
	err := r.db.Raw("UPDATE rooms SET  name = ?, updated_at = ?, type = ? WHERE id IN (SELECT room_id FROM "+
		"rooms r  JOIN  room_users ru ON r.id = ru.room_id WHERE room_id= ? AND user_id = ? )", room.Name, time.Now(),
		room.Type, room.ID, userID).Scan(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepoPG) DeleteRoom(userID, roomID int64) (bool, error) {
	res := r.db.Exec("DELETE FROM rooms WHERE owner_id = ? AND id = ?", userID, roomID)
	err := res.Error
	if err != nil {
		return false, err
	}
	if res.RowsAffected < 1 {
		return false, nil
	}
	return true, err
}

func NewRoomRepo(db *gorm.DB) *RoomRepoPG {
	return &RoomRepoPG{
		db: db,
	}
}
