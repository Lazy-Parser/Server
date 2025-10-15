package token

import (
	"errors"
	"time"

	"github.com/Lazy-Parser/Server/entity"
	"github.com/golang-jwt/jwt/v5"
)

const accessSecret = "secretA"
const AccessTTL = time.Minute * 10
const refreshSecret = "secretB"
const RefreshTTL = time.Hour * 24

type MyClaims struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user entity.User) (string, error) {
	claims := MyClaims{
		user.ID,
		user.RoleID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTTL)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := t.SignedString([]byte(accessSecret))
	return accessToken, err
}

func GenerateRefreshToken(user entity.User) (string, error) {
	claims := MyClaims{
		user.ID,
		user.RoleID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTTL)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := t.SignedString([]byte(refreshSecret))
	return accessToken, err
}

func ParseAccessToken(token string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func ParseRefreshToken(token string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshSecret), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
