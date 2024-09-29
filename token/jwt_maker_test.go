package token

import (
    "testing"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
    maker, err := NewJWTMaker("12345678901234567890123456789012")
    require.NoError(t, err)

    username := "testuser"
    duration := time.Minute

    issuedAt := time.Now()
    expiredAt := issuedAt.Add(duration)

    token, err := maker.CreateToken(username, duration)
    require.NoError(t, err)
    require.NotEmpty(t, token)

    payload, err := maker.VerifyToken(token)
    require.NoError(t, err)
    require.NotEmpty(t, payload)

    require.NotZero(t, payload.ID)
    require.Equal(t, username, payload.Username)
    require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
    require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
    maker, err := NewJWTMaker("12345678901234567890123456789012")
    require.NoError(t, err)

    token, err := maker.CreateToken("testuser", -time.Minute)
    require.NoError(t, err)
    require.NotEmpty(t, token)

    payload, err := maker.VerifyToken(token)
    require.Error(t, err)
    require.EqualError(t, err, "invalid token: token has invalid claims: token is expired")
    require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
    payload, err := NewPayload("testuser", time.Minute)
    require.NoError(t, err)

    jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
        "id":       payload.ID.String(),
        "username": payload.Username,
        "iat":      payload.IssuedAt.Unix(),
        "exp":      payload.ExpiredAt.Unix(),
    })

    token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
    require.NoError(t, err)

    maker, err := NewJWTMaker("12345678901234567890123456789012")
    require.NoError(t, err)

    payload, err = maker.VerifyToken(token)
    require.Error(t, err)
    require.EqualError(t, err, "invalid token: token is unverifiable: error while executing keyfunc: unexpected token signing method")
    require.Nil(t, payload)
}