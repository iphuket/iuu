package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/iuu/app/config"
	"github.com/iphuket/iuu/library/jwt"
)

var (
	aeskey = config.JWTAesKEY
)

// Check user auth
func Check(c *gin.Context) (UserUUID string, err error) {
	UserUUID, err = jwt.Chcek(c.ClientIP(), aeskey, c.GetHeader("iuu_token"))
	return
}

// Login user login
func Login(c *gin.Context) (token string, err error) {
	token, err = jwt.NewToken("uuid", "sub", "ip", aeskey, 60*30)
	// c.SetCookie("iuu_token", token, 60*30, "/", c.Request.URL.Host, true, true)
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
}
