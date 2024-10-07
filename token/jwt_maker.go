package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload,  err
	}

	claims := jwt.MapClaims{
		"id":       payload.ID,
		"username": payload.Username,
		"iat":      payload.IssuedAt.Unix(),
		"exp":      payload.ExpiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(maker.secretKey))
    return accessToken, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
    keyFunc := func(token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
            return nil, fmt.Errorf("unexpected token signing method")
        }
        return []byte(maker.secretKey), nil
    }

    jwtToken, err := jwt.Parse(token, keyFunc)
    if err != nil {
        return nil, fmt.Errorf("invalid token: %w", err)
    }

    claims, ok := jwtToken.Claims.(jwt.MapClaims)
    if !ok {
        return nil, fmt.Errorf("invalid token claims")
    }

    id, err := uuid.Parse(claims["id"].(string))
    if err != nil {
        return nil, fmt.Errorf("invalid token ID")
    }

    username, ok := claims["username"].(string)
    if !ok {
        return nil, fmt.Errorf("invalid username claim")
    }

    issuedAt := time.Unix(int64(claims["iat"].(float64)), 0)
    expiredAt := time.Unix(int64(claims["exp"].(float64)), 0)

    payload := &Payload{
        ID:        id,
        Username:  username,
        IssuedAt:  issuedAt,
        ExpiredAt: expiredAt,
    }

    return payload, nil
}
