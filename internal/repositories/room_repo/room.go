package room_repo

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"time"
)

type RoomType = int8

const (
	PrivateRoom RoomType = iota + 1
	PublicRoom
)

type Room struct {
	ID      int64           `gorm:"column:id;primaryKey"`
	Name    string          `gorm:"column:name"`
	OwnerID int64           `gorm:"column:owner_id;foreignKey,references:ID"`
	Owner   *user_repo.User `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //foreignKey, Belongs To
	Type    int8            `gorm:"column:type"`                                                                    // 1 - private room, 2 - group room
	//RoomUsers []RoomUser      `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Room) TableName() string {
	return "rooms"
}
