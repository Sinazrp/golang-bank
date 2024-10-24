package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sinazrp/golang-bank/api"
	db "github.com/sinazrp/golang-bank/db/sqlc"
	"github.com/sinazrp/golang-bank/gapi"
	"github.com/sinazrp/golang-bank/pb"
	"github.com/sinazrp/golang-bank/util"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	config, err := util.LoadConfig(false)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dbSource := config.DBSource
	connection, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewSQLStore(connection)
	runGrpcServer(store, config)

}

func runGinServer(store db.Store, config util.Config) {
	server, err := api.NewServer(store, config)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

// runGrpcServer starts a gRPC server with the given store and config.
// It sets up the server, registers the server with the gRPC reflection service,
// and begins listening on the TCP address specified in the config.
// If it fails to start the server, it logs the error and exits the program.
func runGrpcServer(store db.Store, config util.Config) {

	server, err := gapi.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGolangBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, errors := net.Listen("tcp", config.GrpcServerAddress)
	if errors != nil {
		log.Fatal("cannot create listener:", errors)
	}
	log.Printf("start Grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}

}
