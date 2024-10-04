package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sinazrp/golang-bank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing for the account
// resource.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Create account
	router.POST("/accounts", server.createAccount)

	// Get account
	router.GET("/accounts/:id", server.getAccount)

	// List accounts
	router.GET("/accounts", server.getAccountList)

	// Delete account
	router.DELETE("/accounts/:id", server.deleteAccount)

	// Update account
	router.PUT("/accounts", server.updateAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse returns a gin.H with an "error" key and the given error message.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
