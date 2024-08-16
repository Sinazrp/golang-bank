package db

import (
	"context"
	"github.com/sinazrp/golang-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomEntry(t *testing.T, account Account) (Entry, error, CreateEntryParams) {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	return entry, err, arg

}
func TestCreateEntry(t *testing.T) {
	account, _, _ := CreateRandomAccount(t)
	entry, err, arg := CreateRandomEntry(t, account)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
}

func TestGetEntry(t *testing.T) {
	account, _, _ := CreateRandomAccount(t)
	entry, err, arg := CreateRandomEntry(t, account)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
}
func TestUpdateEntry(t *testing.T) {
	account, _, _ := CreateRandomAccount(t)
	entry, err, _ := CreateRandomEntry(t, account)

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, account.ID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.NotEqual(t, entry.Amount, entry2.Amount)

}

func TestDeleteEntry(t *testing.T) {
	account, _, _ := CreateRandomAccount(t)
	entry, err, _ := CreateRandomEntry(t, account)
	err = testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	account, _, _ := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, account)
	}
	arq := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arq)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

	// Test with invalid limit
	arq = ListEntriesParams{
		Limit:  -1,
		Offset: 5,
	}
	entries, err = testQueries.ListEntries(context.Background(), arq)
	require.Error(t, err)
	// Test with invalid offset
	arq = ListEntriesParams{
		Limit:  5,
		Offset: -1,
	}
	entries, err = testQueries.ListEntries(context.Background(), arq)
	require.Error(t, err)

}
