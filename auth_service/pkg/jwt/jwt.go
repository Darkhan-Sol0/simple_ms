package jwt

import (
	"auth_service/infrastructure/config"
	"auth_service/internal/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UUID  string
	Login string
	Role  string
	jwt.RegisteredClaims
}

func GenerateToken(user dto.DtoUserToToken) (string, error) {
	claims := &Claims{
		UUID:  user.UUID,
		Login: user.Login,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GetJwtEnv().TokenLifetime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.GetJwtEnv().JWTKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtEnv().JWTKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
