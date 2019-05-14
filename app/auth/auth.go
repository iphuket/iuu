package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/config"
	"github.com/iphuket/pkt/library/crypto"
	"github.com/iphuket/pkt/library/jwt"
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
	eninfo, err := jwt.Chcek(secret, token, c.Request.Host)
	userid = eninfo.UserID
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
	pa := jwt.Payload{
		Issuer:         jp.Issuer,
		Subject:        jp.Subject,
		Audience:       jp.Audience,
		ExpirationTime: now.Add(time.Minute * 30).Unix(),
		NotBefore:      now.Unix(),
		IssuedAt:       now.Unix(),
		JWTID:          "Non-existent",
		EnInfo: jwt.EnInfo{
			UserID: userid,
			IP:     ip,
		},
	}
	token, err = jwt.NewToken(pa, secret)
	return
}

// Renewal user states
func Renewal(c *gin.Context) (err error) {
	token, err := c.Cookie("token")
	if err != nil {
		return err
	}
	c.SetCookie("token", token, 60*30, "/", c.Request.URL.Host, true, true)
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
