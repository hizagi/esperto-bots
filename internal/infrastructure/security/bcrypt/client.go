package bcrypt

import "golang.org/x/crypto/bcrypt"

type BcryptClient struct{}

func NewBcryptClient() *BcryptClient {
	return &BcryptClient{}
}

func (client *BcryptClient) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (client *BcryptClient) Verify(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
