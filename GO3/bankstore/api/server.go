package api

import (
	db "github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/db/sqlc"
	"github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config utils.Config, store *db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.PUT("/accounts", server.updateAccountBalance)
	router.PATCH("/accounts", server.updateAccountOwner)
	//TODO: update owner
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:id", server.deleteAccount)

	router.GET("/accounts/:id/entries", server.listAccountEntries)
	router.GET("/accounts/:id/entries/:entry_id", server.getEntry)
	router.POST("/accounts/:id/entries/", server.createEntry)

	router.GET("/accounts/:id/transfers", server.listTransfers)
	router.GET("/accounts/:id/transfers/:entry_id", server.getTransfer)
	router.POST("/accounts/:id/transfers/", server.createTransfer)

	router.POST("/accounts/:id/sendmoney", server.sendMoney)
	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
