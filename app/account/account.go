package account

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/auth"
	"github.com/iphuket/pkt/app/config"
	"github.com/iphuket/pkt/server"
)

// Route account
func Route(r *gin.RouterGroup) {
	lr := r.Group("/")
	lr.Use(account)
	lr.StaticFile("login", "./static/page/account/login.html")
	lr.StaticFile("register", "./static/page/account/register.html")
	r.StaticFile("home", "./static/page/account/home.html")
	// r.StaticFile("logout", "./static/page/account/logout.html").Use(account)
}
func account(c *gin.Context) {
	_, err := auth.Check(c, server.RemoteIP(c.Request))
	if err != nil {
		// c.JSON(c.Writer.Status(), gin.H{"errCode": "error", "info": fmt.Sprint(err)})
		return
	}
	// c.JSON(c.Writer.Status(), gin.H{"errCode": "success", "uuid": userid, "info": "你已经登录了!"})
	c.Redirect(307, config.SiteConfig().Home)
	c.Abort()
}
