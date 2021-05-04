package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/thanhftu/simple_bank/db/sqlc"
)

// Server serves http request for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a new http server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/api/users", server.createUser)

	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.GET("/api/accounts", server.listAccounts)

	router.POST("/api/transfers", server.createTransfer)
	server.router = router
	return server
}
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
