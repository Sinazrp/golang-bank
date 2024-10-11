package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sinazrp/golang-bank/db/sqlc"
	"github.com/sinazrp/golang-bank/util"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func NewTestServer(store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server")
	}
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}
