package server

import (
	"github.com/gin-gonic/gin"
)

// Engine GIN
var Engine = gin.Default()

// init auto middleware
func init() {
	Engine.Use(server)
	Engine.GET("remoteip", remoteIP)
}

func remoteIP(c *gin.Context) {
	c.JSON(200, gin.H{"RemoteIP": RemoteIP(c.Request)})
}

// server ...
func server(c *gin.Context) {
	c.Header("Server", "iphuket/v1")
}
