package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"os"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
