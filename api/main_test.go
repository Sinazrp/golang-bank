package api

import (
	"github.com/gin-gonic/gin"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}
