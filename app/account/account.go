// Package account ... url 参数 c.Request.FormValue("oc") 即为来源网址
package account

import (
	"fmt"

	"github.com/iphuket/pkt/app/config"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/auth"
	"github.com/iphuket/pkt/server"
)

// Route account
func Route(r *gin.RouterGroup) {
	lr := r.Group("/")
	lr.Use(Account)
	lr.StaticFile("login", "./templates/page/account/login.html")
	lr.StaticFile("register", "./templates/page/account/register.html")

	r.GET("logout", auth.Logout)
	r.StaticFile("reset/main", "./templates/page/account/reset/main.html")
	r.StaticFile("reset/passwd", "./templates/page/account/reset/passwd.html")
	r.StaticFile("home", "./templates/page/account/home.html")

	// r.StaticFile("logout", "./templates/page/account/logout.html").Use(account)
}

// Account M
func Account(c *gin.Context) {
	_, err := auth.Check(c, server.RemoteIP(c.Request))
	if err != nil {
		fmt.Println("account err: ", err)
		return
	}
	if len(c.Request.FormValue("oc")) > 0 {
		c.Redirect(307, c.Request.FormValue("oc"))
		c.Abort()
		return
	}
	c.Redirect(307, config.SiteConfig().Domain)
	c.Abort()
	return
}
