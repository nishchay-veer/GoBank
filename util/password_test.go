package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword (t *testing.T) {
	// Test HashPassword
	password := "password"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)
	
	// Test CheckPassword
	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := "wrong_password"
	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error()) 


}