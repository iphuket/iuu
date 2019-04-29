package server

import (
	"github.com/gin-gonic/gin"
)

var engine = gin.New()

// init auto middleware
func init() {
	engine.Use(server)
}

// New server engine is gin
func New() *gin.Engine {
	return engine
}

// server ...
func server(c *gin.Context) {
	c.Header("Server", "iphuket/v1")
}
