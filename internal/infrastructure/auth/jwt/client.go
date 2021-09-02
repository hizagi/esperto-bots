package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
)

type JwtClient struct {
	configuration viper.Configuration
}

func NewJwtClient(configuration viper.Configuration) *JwtClient {
	return &JwtClient{
		configuration,
	}
}

func (client *JwtClient) CreateToken(user entities.User, expiresIn int64) (string, error) {
	metadata := jwt.MapClaims{}
	metadata["authorized"] = true
	metadata["exp"] = expiresIn
	metadata["userID"] = user.ID
	metadata["userEmail"] = user.Email
	metadata["roles"] = user.Roles

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, metadata)

	return token.SignedString([]byte(client.configuration.AuthConfiguration.SecretKey))
}
