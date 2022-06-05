package user

type AuthService interface {
	SingIn(email, password string) (*User, error) // bcrypt
	SignUp(user User, credentials UserCredentials) (*User, error)
	UpdatePassword(id int64, password string) (bool, error)
}

type AuthServiceImpl struct {
	AuthService AuthService
}

func (a AuthServiceImpl) SingIn(email, password string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) SignUp(user User, credentials UserCredentials) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) UpdatePassword(id int64, password string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthServiceImpl(authService AuthService) *AuthServiceImpl {
	return &AuthServiceImpl{AuthService: authService}
}
