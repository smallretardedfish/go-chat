package room_repo

type RoomUserStatusType = int8

const (
	RoomUserCreated RoomUserStatusType = iota + 1
	RoomUserDeleted
)

type RoomUser struct {
	RoomID int64 `gorm:"column:room_id;primaryKey,foreignKey,references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID int64 `gorm:"column:user_id;primaryKey,foreignKey,references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status int8  `gorm:"column:status"` // 1 - created, 2 - deleted
}

func (RoomUser) TableName() string {
	return "room_users"
}
