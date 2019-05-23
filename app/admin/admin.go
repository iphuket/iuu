package admin

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iphuket/pkt/app/auth"
	"github.com/iphuket/pkt/app/config"
	"github.com/iphuket/pkt/library/crypto"
	"github.com/iphuket/pkt/library/passwd"
	"github.com/iphuket/pkt/server"
)

var (
	jwtKey    = config.JWTSecret
	passwdKey = config.PasswdSecret
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
	jp := auth.JWTPayload{
		Issuer:   mu.Name,
		Subject:  mu.Privilege,
		Audience: []string{mu.Audience},
		UserID:   mu.UUID,
		IP:       server.RemoteIP(c.Request),
	}
	token, err := auth.Token(c, jp)
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	c.SetCookie("token", string(token), 60*30, "/", server.RemoteIP(c.Request), false, false)
	successHandle(c, string(token))
}

// Logout admin
func Logout(c *gin.Context) {
	// 清除缓存
	auth.Logout(c)
}

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

// Check 检查用户是否登陆
func Check(c *gin.Context) {
	uuid, err := auth.Check(c, server.RemoteIP(c.Request))
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	successHandle(c, uuid)
}

// SendResetToken 发送验证码Token 暂时仅支持邮件重置密码
func SendResetToken(c *gin.Context) {
	// var mu ManageUser

}

// StartResetPasswd 开始密码重置 账户 接收 验证方法 method 内容 vc 令牌 token 密码 passwd
func StartResetPasswd(c *gin.Context) {
	err := verify(c.Request.FormValue("method"), c.Request.FormValue("vc"), c.Request.FormValue("token"))
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	// update passwd and code
	var mu ManageUser
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
	err = db.Model(&mu).Where(c.Request.FormValue("method")+" = ?", c.Request.FormValue("vc")).Update(&ManageUser{Passwd: passwd, Code: code}).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	successHandle(c, "ok")
}

// verify 验证
func verify(reMethod, vc, token string) error {
	var mr ManageResets
	db, err := config.DB()
	if err != nil {
		return err
	}
	err = db.Where(reMethod+" = ?", vc).Find(&mr).Error
	if err != nil {
		return err
	}
	if token == mr.Token {
		return nil
	}
	return errors.New("token error")
}
