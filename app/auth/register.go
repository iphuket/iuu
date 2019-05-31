package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iphuket/pkt/app/config"
	"github.com/iphuket/pkt/library/crypto"
	"github.com/iphuket/pkt/library/passwd"
	"github.com/iphuket/pkt/server"
)

// Register admin 注册管理用户
func Register(c *gin.Context) {
	db, err := config.DB()
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	code := crypto.RandomString(6) // 第一次生成 ，第二次用以效果验证
	passwd, err := passwd.New(c.Request.FormValue("passwd"), code)
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	// AutoMigrate sql
	err = db.AutoMigrate(ManageUser{}).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	err = db.Create(&ManageUser{
		UUID:      uuid.New().String(),
		Name:      c.Request.FormValue("name"),
		NickName:  c.Request.FormValue("nickname"),
		RealName:  c.Request.FormValue("realname"),
		Email:     c.Request.FormValue("email"),
		Tel:       c.Request.FormValue("tel"),
		IDCard:    c.Request.FormValue("idcard"),
		CardClass: c.Request.FormValue("cardclass"),
		Passwd:    passwd,
		Code:      code,
		IP:        server.RemoteIP(c.Request),
	}).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	successHandle(c, "ok")
}
