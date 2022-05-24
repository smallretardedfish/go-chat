package message_repo

import (
	"os/user"
	"time"
)

type Message struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	Text      string    `gorm:"column:text"`
	OwnerID   int64     `gorm:"column:owner_id"`
	Owner     user.User `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //foreignKey, Belongs To
	RoomID    int64     `gorm:"column:room_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Message) TableName() string {
	return "messages"
}