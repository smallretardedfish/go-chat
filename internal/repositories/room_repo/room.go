package room_repo

import (
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"time"
)

type Room struct {
	ID        int64            `gorm:"column:id;primaryKey"`
	Name      string           `gorm:"column:name"`
	OwnerID   int64            `gorm:"column:owner_id;foreignKey,references:ID"`
	Owner     user_repo.User   `gorm:"associationForeignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //foreignKey, Belongs To
	Type      int8             `gorm:"column:type"`                                                                 // 1 - private room, 2 - group room
	Users     []user_repo.User `gorm:"many2many:user_rooms;"`                                                       //TODO use many2many
	CreatedAt time.Time        `gorm:"column:created_at"`
	UpdatedAt time.Time        `gorm:"column:updated_at"`
}

func (Room) TableName() string {
	return "rooms"
}
