package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/config"
	"github.com/iphuket/pkt/library/passwd"
	"github.com/iphuket/pkt/server"
)

// Login admin
func Login(c *gin.Context) {
	// 获取 email 查询sql 获得 code, passwd(MD5)->比对 passwd
	email := c.Request.FormValue("email")
	db, err := config.DB()
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	var mu ManageUser
	// 查询emai
	err = db.Where("email = ?", email).First(&mu).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	// 输入密码验证
	passwdIn, err := passwd.New(c.Request.FormValue("passwd"), mu.Code)
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	// 密码验证
	if passwdIn != mu.Passwd {
		errorHandle(c, "error", "passwd error")
		return
	}
	// 验证完成 发出token
	jp := JWTPayload{
		Issuer:   mu.Name,
		Subject:  mu.Privilege,
		Audience: []string{mu.Audience},
		UserID:   mu.UUID,
		IP:       server.RemoteIP(c.Request),
	}
	token, err := Token(c, jp)
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	c.SetCookie("token", string(token), 60*30, "/", server.RemoteIP(c.Request), false, false)
	successHandle(c, c.Request.FormValue("co"))
}
