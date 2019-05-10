package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/iuu/app/config"
	"github.com/iphuket/iuu/library/jwt"
)

var (
	jwtKey = config.JWTAesKey
)

// Check 检测用户是否登陆
func Check(c *gin.Context, ip string) (UserUUID string, err error) {
	token, err := c.Cookie("iuu_token")
	if err != nil {
		return
	}
	UserUUID, err = jwt.Chcek(ip, jwtKey, token)
	return
}

// NewToken 发放令牌
func NewToken(c *gin.Context, uuid, sub, ip string) (token string, err error) {
	token, err = jwt.NewToken(uuid, sub, ip, jwtKey, 60*30)
	return
}

// Renewal user states
func Renewal(c *gin.Context) (err error) {
	token, err := c.Cookie("iuu_token")
	if err != nil {
		return err
	}
	c.SetCookie("iuu_token", token, 60*30, "/", c.Request.URL.Host, true, true)
	return err
}

// Logout user states
func Logout(c *gin.Context) {
	c.SetCookie("iuu_token", "", -1, "/", c.Request.URL.Host, true, true)
	successHandle(c, "ok")
}

func successHandle(c *gin.Context, info ...interface{}) {
	c.JSON(c.Writer.Status(), gin.H{"errCode": "success", "info": info})
	c.Abort()
}
