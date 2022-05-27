package user_repo

type UserCredentials struct {
	ID       int64  `gorm:"column:id;primaryKey;foreignKey,references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Password string `gorm:"column:password"` // bcrypt
}

func (UserCredentials) TableName() string {
	return "user_credentials"
}
