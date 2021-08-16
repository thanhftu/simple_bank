package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/thanhftu/simple_bank/db/sqlc"
	"github.com/thanhftu/simple_bank/token"
	"github.com/thanhftu/simple_bank/utils"
)

// Server serves http request for our banking service
type Server struct {
	config     utils.Config
	store      db.Store
	TokenMaker token.Maker
	router     *gin.Engine
}

// NewServer create a new http server and setup routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenmaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		store:      store,
		config:     config,
		TokenMaker: tokenmaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/api/users", server.createUser)
	router.POST("/api/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.TokenMaker))

	authRoutes.POST("/api/accounts", server.createAccount)
	authRoutes.GET("/api/accounts/:id", server.getAccount)
	authRoutes.GET("/api/accounts", server.listAccounts)

	authRoutes.POST("/api/transfers", server.createTransfer)
	server.router = router

}
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
