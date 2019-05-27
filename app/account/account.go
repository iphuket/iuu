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
	lr.Use(account)
	lr.StaticFile("login", "./static/page/account/login.html")
	lr.StaticFile("register", "./static/page/account/register.html")

	r.StaticFile("reset/main", "./static/page/account/reset/main.html")
	r.StaticFile("reset/passwd", "./static/page/account/reset/passwd.html")
	r.StaticFile("home", "./static/page/account/home.html")

	// r.StaticFile("logout", "./static/page/account/logout.html").Use(account)
}
func account(c *gin.Context) {
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
	c.Redirect(307, config.SiteConfig().Home)
	c.Abort()
	return
}
