package server

import (
	"github.com/gin-gonic/gin"
)

var engine = gin.Default()

// init auto middleware
func init() {
	engine.Use(server)
	engine.GET("remoteip", remoteIP)
}

func remoteIP(c *gin.Context) {
	c.JSON(200, gin.H{"RemoteIP": RemoteIP(c.Request)})
}

// New server engine is gin
func New() *gin.Engine {
	return engine
}

// server ...
func server(c *gin.Context) {
	c.Header("Server", "iphuket/v1")
}
