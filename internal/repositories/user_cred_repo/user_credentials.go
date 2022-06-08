package user_cred_repo

type UserCredentials struct {
	Email    string `gorm:"column:email;unique"`
	Password string `gorm:"column:password"`
}

func (UserCredentials) TableName() string {
	return "user_credentials"
}
