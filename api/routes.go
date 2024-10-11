package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRoutes() {
	router := gin.Default()
	// Create an account
	router.POST("/accounts", server.createAccount)

	// Get an account
	router.GET("/accounts/:id", server.getAccount)

	// List accounts
	router.GET("/accounts", server.getAccountList)

	// Delete account
	router.DELETE("/accounts/:id", server.deleteAccount)

	// Update account
	router.PUT("/accounts", server.updateAccount)

	// Create Transfer
	router.POST("/transfers", server.createTransfer)

	//create User
	router.POST("/user", server.createUser)

	// Login User
	router.POST("/login", server.LoginUser)
	
	server.router = router

}
