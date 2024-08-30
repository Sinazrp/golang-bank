package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// first start tx and then create new queries and call fn with queries, if no error return tx.Commit()
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
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

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		txName := ctx.Value(txKey)
		fmt.Println(txName, "create Transfer")
		result.Transfer, err = store.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "create entry1")
		result.FromEntry, err = store.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "create entry2")
		result.ToEntry, err = store.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")

		result.FromAccount, err = store.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return nil
		}

		fmt.Println(txName, "update account 2")
		result.ToAccount, err = store.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return nil
		}
		return nil
	})

	return result, err
}