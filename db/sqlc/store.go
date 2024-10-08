package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store Interface: This defines the contract for the store, which includes all the methods of
// the Querier (e.g., basic CRUD operations) and a TransferTx method for handling transactions.
// This interface allows you to have multiple implementations of Store (e.g., SQL-based, NoSQL-based, in-memory) if needed.
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore struct: This is a concrete implementation of the Store interface, and it wraps around
// the SQL database (sql.DB). It extends the functionality provided by Queries with transaction
// handling, via the execTx method and TransferTx.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewSQLStore we can use SQLStore here because we implement the
// Store interface, and it used as receiver in execTx and TransferTx methods
func NewSQLStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// first start tx and then create new queries and call fn with queries, if no error return tx.Commit()
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(queries *Queries) error {
		var err error

		result.Transfer, err = store.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = store.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = store.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		//This if statement is used to determine the order in which the addMoney function is called,
		//depending on the relative values of FromAccountID and ToAccountID. By ordering the accounts
		//consistently (based on their IDs), it helps prevent potential deadlocks when there are multiple
		//simultaneous transactions between the same accounts.

		if arg.FromAccountID < arg.ToAccountID {

			result.FromAccount, result.ToAccount, err = addMoney(ctx, queries, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, queries, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)

		}
		return nil
	})

	return result, err
}

// addMoney performs two AddAccountBalance operations in a single transaction. It takes
// four parameters: two account IDs and two amounts. It will add amount1 to the account
// with ID accountID1 and add amount2 to the account with ID accountID2. The function
// returns the two modified accounts and an error (if any).
func addMoney(ctx context.Context, queries *Queries, accountID1 int64, amount1 int64, accountID2 int64, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = queries.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}
	account2, err = queries.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	if err != nil {
		return
	}
	return account1, account2, err

}
