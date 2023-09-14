package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP request
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates an instance of a new HTTP server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	defRouter := gin.Default()

	//defRouter.GET("/accounts", server.GetAccounts)
	defRouter.POST("/account", server.CreateAccount)
	defRouter.GET("/account/:id", server.GetAccountFromId)
	defRouter.GET("/accounts", server.ListAccounts)

	server.router = defRouter
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error(), "message": "Could not complete request."}
}
