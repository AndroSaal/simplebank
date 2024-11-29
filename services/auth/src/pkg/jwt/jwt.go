package jwtAuth

import (
	"time"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
	"github.com/golang-jwt/jwt/v5"
)

const SecretAuthSalt = "OKJSNDOICSDIOCMWPEOCKM987654"

func NewToken(user models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(SecretAuthSalt))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
