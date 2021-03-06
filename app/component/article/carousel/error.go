package carousel

import (
	"github.com/gin-gonic/gin"
)

const (
	// dbConErr 数据库 连接错误
	dbConErr = "Database Connection error "
)

func errorHandle(c *gin.Context, errCode, Info string) {
	c.JSON(c.Writer.Status(), gin.H{"errCode": errCode, "info": Info})
	c.Abort()
}
func successHandle(c *gin.Context, info ...interface{}) {
	c.JSON(c.Writer.Status(), gin.H{"errCode": "success", "info": info})
	c.Abort()
}
