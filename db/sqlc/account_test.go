package db

import (
	"context"
	"github.com/sinazrp/golang-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	arg := util.RandomAccount()

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}
