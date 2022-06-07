package room_repo

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type RoomRepo interface {
	GetRoom(userID, roomID int64) (*Room, error)
	GetRooms(userID int64) ([]Room, error)
	CreateRoom(room Room) (*Room, error)
	CreateRoomUser(roomUser RoomUser) (*RoomUser, error)
	DeleteRoomUser(roomID, userID int64) (bool, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) (bool, error)
}

type RoomRepoPG struct {
	db *gorm.DB
}

func (r *RoomRepoPG) GetRoom(userID, roomID int64) (*Room, error) {
	room := Room{}
	err := r.db.Raw("SELECT * FROM rooms AS r JOIN room_users AS ru ON r.id = ru.room_id WHERE ru.user_id = ? AND ru.room_id = ?", userID, roomID).
		First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepoPG) GetRooms(userID int64) ([]Room, error) { // TODO fix gorm belongs-to population of Users slice
	var rooms []Room
	err := r.db.Raw("SELECT * FROM rooms AS r JOIN room_users AS ru ON r.id = ru.room_id WHERE ru.user_id = ?", userID).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepoPG) CreateRoomUser(roomUser RoomUser) (*RoomUser, error) {
	err := r.db.Raw("INSERT INTO room_users(user_id,room_id) VALUES(?,?) RETURNING *", roomUser.UserID, roomUser.RoomID).Scan(&roomUser).Error
	if err != nil {
		return nil, err
	}
	return &roomUser, nil
}

func (r *RoomRepoPG) DeleteRoomUser(roomID, userID int64) (bool, error) {
	err := r.db.Exec("DELETE FROM room_users WHERE user_id = ? AND room_id = ? ", userID, roomID).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RoomRepoPG) CreateRoom(room Room) (*Room, error) {
	err := r.db.Raw(`INSERT INTO rooms(owner_id,name) VALUES(?,?) RETURNING *`, room.OwnerID, room.Name).Scan(&room).Error
	if err != nil {
		return nil, err
	}
	roomUser := RoomUser{
		RoomID: room.ID,
		UserID: room.OwnerID,
	}
	if _, err := r.CreateRoomUser(roomUser); err != nil {
		return nil, err
	}
	return &room, err
}

func (r *RoomRepoPG) UpdateRoom(userID int64, room Room) (*Room, error) {
	log.Println("UPDATING ROOM WITH ID", room.ID)
	err := r.db.Raw("UPDATE rooms SET  name = ?, updated_at = ? WHERE id IN (SELECT room_id FROM "+
		"rooms r  JOIN  room_users ru ON r.id = ru.room_id WHERE room_id= ? AND user_id = ? )", room.Name, time.Now(), room.ID, userID).Scan(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepoPG) DeleteRoom(userID, roomID int64) (bool, error) {
	err := r.db.Exec("DELETE FROM room_users WHERE owner_id = ? AND room_id = ?", userID, roomID).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func NewRoomRepo(db *gorm.DB) *RoomRepoPG {
	return &RoomRepoPG{
		db: db,
	}
}
