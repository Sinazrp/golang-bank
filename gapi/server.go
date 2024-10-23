package gapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/sinazrp/golang-bank/db/sqlc"
	"github.com/sinazrp/golang-bank/pb"
	"github.com/sinazrp/golang-bank/token"
	"github.com/sinazrp/golang-bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new grpc server.
// resource.
func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err)

	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	return server, nil
}
