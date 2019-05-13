package psutil

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/component/psutil/api"
)

// Route Init
func Route(r *gin.RouterGroup) {
	api.Route(r)
}
