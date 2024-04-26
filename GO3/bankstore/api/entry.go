package api

import (
	"log"
	"net/http"

	db "github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type CreateEntryRequest struct {
	AccountID int64 `json:"account_id" binding:"required,min=1"`
	Amount    int64 `json:"amount" binding:"required,min=0"`
}
type GetEntryRequest struct {
	AccountID int64 `uri:"id" binding:"required,min=1"`
	EntryID   int64 `uri:"entry_id" binding:"required,min=1"`
}
type ListEntriesRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
	// PageID   int32 `query:"page_id" binding:"min=1"`
	// PageSize int32 `query:"page_size" binding:"min=5,max=15"`
}

func (server *Server) createEntry(ctx *gin.Context) {
	var req CreateEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateEntryParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}
	entry, err := server.store.CreateEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entry)
}

func (server *Server) getEntry(ctx *gin.Context) {

	var req GetEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	entry, err := server.store.GetEntry(ctx, req.EntryID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if entry.AccountID != req.AccountID {
		ctx.JSON(http.StatusUnauthorized, struct {
			Message string `json:"message"`
		}{
			Message: "unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, entry)
}

func (server *Server) listAccountEntries(ctx *gin.Context) {
	var req ListEntriesRequest
	ctx.BindUri(&req)
	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Println("2", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListEntriesParams{
		AccountID: req.ID,
		Limit:     1000, //req.PageSize,
		Offset:    0,    //(req.PageID - 1) * req.PageSize,
	}
	entries, err := server.store.ListEntries(ctx, arg)

	log.Println(entries)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entries)
}
