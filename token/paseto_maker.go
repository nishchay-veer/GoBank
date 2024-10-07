package token

import (
    "fmt"
    "time"

    "github.com/o1egl/paseto"
    "golang.org/x/crypto/chacha20poly1305"
)

// PASETOMaker is a PASETO token maker
type PASETOMaker struct {
    paseto       *paseto.V2
    symmetricKey []byte
}

// NewPASETOMaker creates a new PASETOMaker
func NewPASETOMaker(symmetricKey string) (Maker, error) {
    if len(symmetricKey) != chacha20poly1305.KeySize {
        return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
    }

    maker := &PASETOMaker{
        paseto:       paseto.NewV2(),
        symmetricKey: []byte(symmetricKey),
    }

    return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PASETOMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
    payload, err := NewPayload(username, duration)
    if err != nil {
        return "", payload, err
    }

    token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
    return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *PASETOMaker) VerifyToken(token string) (*Payload, error) {
    payload := &Payload{}

    err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
    if err != nil {
        return nil, fmt.Errorf("invalid token: %w", err)
    }

    err = payload.Valid()
    if err != nil {
        return nil, fmt.Errorf("invalid token: %w", err)
    }

    return payload, nil
}