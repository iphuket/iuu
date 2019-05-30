// Package account ... url 参数 c.Request.FormValue("oc") 即为来源网址
package account

import (
	"net/http"

	"github.com/iphuket/pkt/app/admin"

	"github.com/iphuket/pkt/app/config"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/auth"
	"github.com/iphuket/pkt/server"
)

// Route account
func Route(r *gin.RouterGroup) {
	lr := r.Group("/")
	lr.Use(authM)
	// lr.StaticFile("login", "./templates/page/account/login.html")
	// lr.StaticFile("register", "./templates/page/account/register.html")
	// 检查登录以去向
	lr.Any("login", Login)
	lr.Any("register", Register)
	// 不必检查登录
	r.GET("logout", Logout)

	r.StaticFile("reset/main", "./templates/page/account/reset/main.html")
	r.StaticFile("reset/passwd", "./templates/page/account/reset/passwd.html")
	r.StaticFile("home", "./templates/page/account/home.html")

}

// Account M
func authM(c *gin.Context) {
	_, err := auth.Check(c, server.RemoteIP(c.Request))
	if err != nil {
		return
	}
	c.Redirect(307, config.SiteConfig().Domain)
	c.Abort()
	return
}

// Login 登录
func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "page/account/login.html", gin.H{
			"lang":   "zh",
			"title":  "pkt 登录页面",
			"logout": config.SiteConfig().Login + "?co=" + c.Request.URL.String(),
			"api":    gin.H{"create": "/component/shoturl/api?do=create", "delete": "/component/shoturl/api?do=delete"},
		})
	}
	if c.Request.Method == "POST" {
		admin.Login(c)
	}
}

// Register 注册
func Register(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "page/account/register.html", gin.H{
			"lang":   "zh",
			"title":  "短网址管理页",
			"logout": config.SiteConfig().Login + "?co=" + c.Request.URL.String(),
			"api":    gin.H{"create": "/component/shoturl/api?do=create", "delete": "/component/shoturl/api?do=delete"},
		})
	}
	if c.Request.Method == "POST" {
		admin.Register(c)
	}
}

// Logout 退出登录
func Logout(c *gin.Context) {
	auth.Logout(c)
}
