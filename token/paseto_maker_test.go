package token

import (
    "testing"
    "time"

    "github.com/stretchr/testify/require"
)

func TestPASETOMaker(t *testing.T) {
    maker, err := NewPASETOMaker("12345678901234567890123456789012")
    require.NoError(t, err)

    username := "testuser"
    duration := time.Minute

    issuedAt := time.Now()
    expiredAt := issuedAt.Add(duration)

    token, Payload, err := maker.CreateToken(username, duration)
    require.NoError(t, err)
    require.NotEmpty(t, token)
    require.NotEmpty(t, Payload)

    payload, err := maker.VerifyToken(token)
    require.NoError(t, err)
    require.NotEmpty(t, payload)

    require.NotZero(t, payload.ID)
    require.Equal(t, username, payload.Username)
    require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
    require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPASETOToken(t *testing.T) {
    maker, err := NewPASETOMaker("12345678901234567890123456789012")
    require.NoError(t, err)

    token, Payload, err := maker.CreateToken("testuser", -time.Minute)
    require.NoError(t, err)
    require.NotEmpty(t, token)
    require.NotEmpty(t, Payload)

    payload, err := maker.VerifyToken(token)
    require.Error(t, err)
    require.EqualError(t, err, "invalid token: token has expired")
    require.Nil(t, payload)
}