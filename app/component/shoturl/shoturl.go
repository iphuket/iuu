// Package shoturl ...doc
// domain/api/shoturl
// fromdata:
// ?do=create&source=要缩短的网址;
// ?do=delete&str=要删除的短网址;
package shoturl

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/iphuket/pkt/app/account"

	"github.com/google/uuid"

	"github.com/iphuket/pkt/app/auth"

	"github.com/iphuket/pkt/server"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/config"
)

// Route Init
func Route(en *gin.Engine) {
	en.Any("s/:code", autoRedirect)

	en.LoadHTMLGlob("templates/**/*/*")
	st := en.Group("component/shoturl")
	st.Use(account.Account)
	st.Any("api", control)
	st.GET("main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "component/shoturl/main.html", gin.H{
			"lang":   "zh",
			"title":  "短网址管理页",
			"logout": config.SiteConfig().Logout + "?co=" + c.Request.URL.String(),
		})
	})

}
func authM(c *gin.Context) {
	_, err := auth.Check(c, server.RemoteIP(c.Request))
	if err != nil {
		c.Redirect(307, config.SiteConfig().Login+"?co="+c.Request.URL.String())
		c.Abort()
		return
	}
	c.Next()
}

var domain = config.ShotURLDomain

// autoRedirect Redirect
func autoRedirect(c *gin.Context) {
	db, err := config.DB()
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	var sURL ShotURL
	//
	err = db.Where("code = ?", c.Param("code")).First(&sURL).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	c.Redirect(307, sURL.Protocol+""+sURL.Source)
}

// 管理类型
func control(c *gin.Context) {
	do := c.Request.FormValue("do")
	// 权限验证
	userid, err := auth.Check(c, server.RemoteIP(c.Request))
	if err != nil {
		c.Redirect(307, config.SiteConfig().Login)
		c.Abort()
		return
	}
	switch do {
	case "create":
		create(c, userid)
	case "delete":
		delete(c, userid)
	}
}

// create 创建短网址
func create(c *gin.Context, userid string) {

	db, err := config.DB()
	if err != nil {
		errorHandle(c, "db con ", fmt.Sprint(err))
		return
	}
	var sURL ShotURL
	db.AutoMigrate(sURL)
	int, err := strconv.Atoi(c.Request.FormValue("length"))
	if err != nil {
		errorHandle(c, "string to int error ", fmt.Sprint(err))
		return
	}
	sURL = ShotURL{
		UUID:     uuid.New().String(),
		UserUUID: userid,
		Code:     romStr(int),
		Source:   c.Request.FormValue("source"),
	}
	err = db.Create(&sURL).Error
	if err != nil {
		errorHandle(c, "db Create", fmt.Sprint(err))
		return
	}
	successHandle(c, domain+""+sURL.Code)
}

// delete 删除短网址
func delete(c *gin.Context, userid string) {
	var sURL ShotURL
	db, err := config.DB()
	if err != nil {
		errorHandle(c, "db con ", fmt.Sprint(err))
	}
	err = db.Where("code = ? AND user_uuid = ?", c.Param("code"), userid).Delete(sURL).Error
	if err != nil {
		errorHandle(c, "db where ", fmt.Sprint(err))
	}
	successHandle(c, "delete "+c.Param("code")+" success")
}

// 随机字符串 ...
func romStr(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
