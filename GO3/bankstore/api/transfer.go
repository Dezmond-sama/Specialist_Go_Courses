package api

import (
	"net/http"

	db "github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type GetTransferRequest struct {
	TransferID int64 `uri:"id" binding:"required,min=1"`
	AccountID  int64 `uri:"account_id" binding:"required,min=1"`
}

type ListTransfersRequest struct {
	AccountID int64 `uri:"id" binding:"required,min=1"`
	// PageID   int32 `query:"page_id" binding:"min=1"`
	// PageSize int32 `query:"page_size" binding:"min=5,max=15"`
}

type CreateTransferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64 `json:"to_account_id" binding:"required,min=1"`
	Amount        int64 `json:"amount" binding:"required,gt=0"`
}

func (server *Server) getTransfer(ctx *gin.Context) {

	var req GetTransferRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	transfer, err := server.store.GetTransfer(ctx, req.TransferID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if transfer.FromAccountID != req.AccountID && transfer.ToAccountID != req.AccountID {
		ctx.JSON(http.StatusUnauthorized, struct {
			Message string `json:"message"`
		}{
			Message: "unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, transfer)
}

func (server *Server) listTransfers(ctx *gin.Context) {
	var req ListTransfersRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListTransfersForAccountParams{
		FromAccountID: req.AccountID,
		Limit:         1000,
		Offset:        0,
	}
	transfers, err := server.store.ListTransfersForAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transfers)
}

func (server *Server) createTransfer(ctx *gin.Context) {

	var req CreateTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	transfer, err := server.store.CreateTransfer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transfer)
}
