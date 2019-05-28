package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/config"
	"github.com/iphuket/pkt/library/crypto"
	"github.com/iphuket/pkt/library/jwt"
	"github.com/iphuket/pkt/server"
)

var (
	secret = config.JWTSecret
)

// Check 检测用户是否登陆 if len(nowip) > 0  check nowip == lasip 则验证通过
func Check(c *gin.Context, nowip string) (userid string, err error) {
	token, err := c.Cookie("token")
	userid = ""
	if err != nil {
		return userid, err
	}
	eninfo, err := jwt.Chcek(secret, token, server.RemoteIP(c.Request))
	if err != nil {
		return userid, err
	}
	userid, err = crypto.DeCrypt(eninfo.UserID, []byte(secret))
	if err != nil {
		return userid, err
	}
	if len(nowip) > 0 {
		lastip, err := crypto.DeCrypt(eninfo.IP, []byte(secret))
		if err != nil {
			return userid, err
		}
		err = jwt.IPChcek(nowip, lastip)
		return userid, err
	}
	return userid, err
}

// JWTPayload ...
type JWTPayload struct {
	Issuer   string
	Subject  string
	Audience []string
	UserID   string
	IP       string
}

// Token 发放令牌
func Token(c *gin.Context, jp JWTPayload) (token []byte, err error) {
	now := time.Now()
	userid, err := crypto.Encrypt([]byte(jp.UserID), []byte(secret))
	if err != nil {
		return token, err
	}
	ip, err := crypto.Encrypt([]byte(jp.IP), []byte(secret))
	if err != nil {
		return token, err
	}
	// pa := jwt.Payload{IssuedAt: "Ss"}

	p := new(jwt.Payload)
	p.Issuer = jp.Issuer
	p.Subject = jp.Subject
	p.Audience = jp.Audience
	p.ExpirationTime = now.Add(time.Minute * 30).Unix()
	p.NotBefore = now.Unix()
	p.IssuedAt = now.Unix()
	p.JWTID = "Non-existent"
	p.EnInfo.UserID = userid
	p.EnInfo.IP = ip
	token, err = jwt.NewToken(p, secret)
	return
}

// Renewal user states
func Renewal(c *gin.Context) (err error) {
	token, err := c.Cookie("token")
	if err != nil {
		return err
	}
	fmt.Println(server.RemoteIP(c.Request))
	c.SetCookie("token", token, 60*30, "/", server.RemoteIP(c.Request), false, false)
	return err
}

// Logout user states
func Logout(c *gin.Context) {
	c.SetCookie("token", "logout", -1, "/", server.RemoteIP(c.Request), false, false)
	//successHandle(c, "logout")
	c.Redirect(307, config.SiteConfig().Login+"?co="+c.Request.URL.String())
	c.Abort()
}

func successHandle(c *gin.Context, info ...interface{}) {
	c.JSON(c.Writer.Status(), gin.H{"errCode": "success", "info": info})
	c.Abort()
}
