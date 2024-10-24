package gapi

import (
	"context"
	"fmt"
	db "github.com/sinazrp/golang-bank/db/sqlc"
	"github.com/sinazrp/golang-bank/pb"
	"github.com/sinazrp/golang-bank/token"
	"github.com/sinazrp/golang-bank/util"
)

type Server struct {
	pb.UnimplementedGolangBankServer
	config     util.Config
	store      db.Store
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
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// implementation here
	return nil, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// implementation here
	return nil, nil
}
