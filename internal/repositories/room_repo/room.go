package room_repo

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"time"
)

type Room struct {
	ID        int64          `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	OwnerID   int64          `gorm:"column:owner_id"`
	Owner     user_repo.User `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //foreignKey, Belongs To
	Type      int8           `gorm:"type"`                                                                           // 1 - private room, 2 - group room
	RoomUsers []RoomUser     `gorm:"foreignkey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (Room) TableName() string {
	return "rooms"
}
