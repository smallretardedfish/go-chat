package user_repo

type UserCredentials struct {
	ID       int64  `gorm:"column:id;primaryKey"`
	Password string `gorm:"column:password"` // bcrypt
}
