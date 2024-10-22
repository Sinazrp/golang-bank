package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRoutes() {

	router := gin.Default()
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// Create an account
	authRoutes.POST("/accounts", server.createAccount)

	// Get an account
	authRoutes.GET("/accounts/:id", server.getAccount)

	// List accounts
	authRoutes.GET("/accounts", server.getAccountList)

	// Delete account
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	// Update account
	authRoutes.PUT("/accounts", server.updateAccount)

	// Create Transfer
	authRoutes.POST("/transfers", server.createTransfer)

	//create User
	router.POST("/user", server.createUser)

	// Login User
	router.POST("/login", server.LoginUser)
	//get User
	authRoutes.GET("/user", server.getUser)

	router.POST("/token/renew_access", server.RenewAccessToken)

	server.router = router

}
