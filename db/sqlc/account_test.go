package db

import (
	"context"
	"database/sql"
	"github.com/sinazrp/golang-bank/util"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func CreateRandomAccount(t *testing.T) (Account, error, CreateAccountParams) {
	arg := RandomAccount()

	account, err := testQueries.CreateAccount(context.Background(), arg)
	return account, err, arg

}
func TestCreateAccount(t *testing.T) {
	account, err, arg := CreateRandomAccount(t)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
func TestGetAccount(t *testing.T) {
	account1, _, _ := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, 0)
}

func TestUpdateAccount(t *testing.T) {
	account1, _, _ := CreateRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.NotEqual(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, 0)
}

func TestDeleteAccount(t *testing.T) {
	account1, _, _ := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}
func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, err, _ := CreateRandomAccount(t)
		require.NoError(t, err)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	log.Printf("accounts: %v", accounts)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
	// Test with invalid limit
	arg = ListAccountsParams{
		Limit:  -1,
		Offset: 5,
	}
	accounts, err = testQueries.ListAccounts(context.Background(), arg)
	require.Error(t, err)

	// Test with invalid offset
	arg = ListAccountsParams{
		Limit:  5,
		Offset: -1,
	}
	accounts, err = testQueries.ListAccounts(context.Background(), arg)
	require.Error(t, err)

	// Test with large offset
	arg = ListAccountsParams{
		Limit:  5,
		Offset: 1000,
	}
	accounts, err = testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, accounts)

}

func TestAddBalanceAccount(t *testing.T) {
	account1, _, _ := CreateRandomAccount(t)
	arg := AddAccountBalanceParams{
		ID:     account1.ID,
		Amount: util.RandomMoney(),
	}
	account2, err := testQueries.AddAccountBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.NotEqual(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, 0)
}

func TestGetAccountForUpdate(t *testing.T) {
	account1, _, _ := CreateRandomAccount(t)
	account2, err := testQueries.GetAccountForUpdate(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, 0)
}
