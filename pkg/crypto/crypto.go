package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type AppCrypto struct {
}

func (a *AppCrypto) HashAndSalt(password []byte) (string, error) {
	pass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(pass), err
}

func (a *AppCrypto) ComparePasswords(hashedPassword string, password []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
}
