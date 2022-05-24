package message_repo

type UserMessage struct {
	UserID    int64 `gorm:"column:user_id;primaryKey"`
	MessageID int64 `gorm:"column:message_id;primaryKey"`
	Status    int8  `gorm:"column:status"` // 1 - unread, 2 - read
}

func (UserMessage) TableName() string {
	return "user_messages"
}
