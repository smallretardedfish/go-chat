package room_repo

type RoomUser struct {
	RoomID int64 `gorm:"column:room_id;primaryKey,foreignKey,references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID int64 `gorm:"column:user_id;primaryKey,foreignKey,references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (RoomUser) TableName() string {
	return "room_users"
}
