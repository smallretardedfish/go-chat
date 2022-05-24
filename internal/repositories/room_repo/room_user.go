package room_repo

type RoomUser struct {
	RoomID int64 `gorm:"column:room_id;primaryKey"`
	UserID int64 `gorm:"column:user_id;primaryKey"`
	Status int8  `gorm:"column:status"` // 1 - created, 2 - deleted
}

func (RoomUser) TableName() string {
	return "room_users"
}
