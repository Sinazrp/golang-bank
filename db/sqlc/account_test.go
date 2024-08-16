package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	arg := createAccountParams{
		Owner:    "John",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.createAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account2, err := testQueries.getAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}
