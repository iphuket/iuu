package auth

import (
	"github.com/gin-gonic/gin"
)

func respMsg(c *gin.Context, errCode string, info interface{}) {
	c.JSON(c.Writer.Status(), gin.H{"errCode": errCode, "info": info})
	c.Abort()
}

// codeType ... 错误类型
type codeType map[string]interface{}

// errCode 错误消息
var errCode = codeType{
	"80001": "Email Error",
	"80002": "Passwd Error",
	"80003": "Token Error",
	"80004": "Database Connection Error",
}

// ErrCodeDesc 错误描述
func ErrCodeDesc(code string) string {
	return errCode[code].(string)
}
