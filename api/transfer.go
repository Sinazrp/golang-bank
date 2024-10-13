package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/sinazrp/golang-bank/db/sqlc"
	"github.com/sinazrp/golang-bank/token"
	"net/http"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,ID"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,ID"`
	Amount        int64  `json:"amount" binding:"required,amount"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := ctx.MustGet(authorizationPayLoadKey).(*token.Payload)
	if !server.validAccount(ctx, req.FromAccountID, req.Currency, payload, true) {
		return
	}

	if !server.validAccount(ctx, req.ToAccountID, req.Currency, payload, false) {
		return
	}
	if req.FromAccountID == req.ToAccountID {
		err := errors.New("cannot transfer to the same account")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string, payload *token.Payload, checkOwner bool) bool {
	account, err := server.store.GetAccount(ctx, accountID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if checkOwner {
		if account.Owner != payload.Username {
			err := errors.New("account doesn't belong to the authenticated user")
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return false
		}
	}
	if account.Currency != currency {
		err := fmt.Errorf("account doesnt have the neccesary currency should be %s but got %s", currency, account.Currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false

	}

	return true

}
