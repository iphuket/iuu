package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/component/psutil/cpu"
)

// Route Settings
func Route(r *gin.RouterGroup) {
	cp := r.Group("cpu")
	cp.GET("info", cpu.Info)
	cp.GET("usage", cpu.Usage)
	// r.GET("list/:class", list.List)
}
