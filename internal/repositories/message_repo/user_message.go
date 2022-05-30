package message_repo

type UserMessageStatus = int8

const (
	UserMessageUnread UserMessageStatus = iota + 1
	UserMessageRead
)

type UserMessage struct {
	UserID    int64 `gorm:"column:user_id;primaryKey,foreignKey,references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MessageID int64 `gorm:"column:message_id;primaryKey,foreignKey,references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status    int8  `gorm:"column:status"` // 1 - unread, 2 - read
}

func (UserMessage) TableName() string {
	return "user_messages"
}
