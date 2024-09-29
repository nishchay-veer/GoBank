package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nishchay-veer/simplebank/util"
	"github.com/stretchr/testify/require"
)


func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
	
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	Account1 := createRandomAccount(t)
	Account2, err := testQueries.GetAccount(context.Background(), Account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Account2)
	require.Equal(t, Account1.Owner, Account2.Owner)
	require.Equal(t, Account1.Balance, Account2.Balance)
	require.Equal(t, Account1.Currency, Account2.Currency)
	require.Equal(t, Account1.ID, Account2.ID)
	require.WithinDuration(t, Account1.CreatedAt, Account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	Account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      Account1.ID,
		Balance: util.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, Account1.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, Account1.Currency, account.Currency)
	require.Equal(t, Account1.ID, account.ID)
	require.WithinDuration(t, Account1.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	Account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), Account1.ID)
	require.NoError(t, err)
	account, err := testQueries.GetAccount(context.Background(), Account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}


func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListaccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.Listaccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

	

}