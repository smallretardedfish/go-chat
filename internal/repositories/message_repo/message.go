package message_repo

import (
	"github.com/lib/pq"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"time"
)

type Message struct {
	ID           int64          `gorm:"column:id;primaryKey"`
	Text         string         `gorm:"column:text"`
	OwnerID      int64          `gorm:"column:owner_id;foreignKey,references:ID"`
	Owner        user_repo.User `gorm:"associationForeignKey:OwnerID;foreignKey:OwnerID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //foreignKey, Belongs To
	RoomID       int64          `gorm:"column:room_id;foreignKey,references:ID"`
	DeletedUsers pq.Int64Array  `gorm:"column:deleted_users;type:bigint[]; default:'{}'"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

func (Message) TableName() string {
	return "messages"
}
