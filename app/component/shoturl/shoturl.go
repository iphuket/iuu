package shoturl

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/iphuket/pkt/app/auth"

	"github.com/iphuket/pkt/server"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/config"
)

// Route Init
func Route(r *gin.RouterGroup) {
	r.Any("l/:str", lctr)
	r.Any("m", mctr)
}
func authM(c *gin.Context) {
	_, err := auth.Check(c, server.RemoteIP(c.Request))
	if err != nil {
		c.Redirect(307, config.SiteConfig().Login)
		c.Abort()
		return
	}
	c.Next()
}

var domain = config.ShotURLDomain

// lctr 永久短链接转换
func lctr(c *gin.Context) {
	db, err := config.DB()
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	var sURL ShotURL
	// 查询emai
	err = db.Where("str = ?", c.Param("str")).First(&sURL).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	c.Redirect(307, sURL.Protocol+""+sURL.Source)
	//successHandle(c)
}

// tctr 临时短链接
func tctr(c *gin.Context) {
}
func mctr(c *gin.Context) {
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
		Str:      romStr(int),
		Source:   c.Request.FormValue("source"),
		Protocol: c.Request.FormValue("protocol"),
	}
	err = db.Create(&sURL).Error
	if err != nil {
		errorHandle(c, "db Create", fmt.Sprint(err))
		return
	}
	successHandle(c, domain+"/l/"+sURL.Str)
}
func delete(c *gin.Context, userid string) {
	var sURL ShotURL
	db, err := config.DB()
	if err != nil {
		errorHandle(c, "db con ", fmt.Sprint(err))
	}
	err = db.Where("key = ? AND user_uuid = ?", c.Param("str"), userid).Delete(sURL).Error
	if err != nil {
		errorHandle(c, "db where ", fmt.Sprint(err))
	}
	successHandle(c, "delete "+c.Param("str")+" success")
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
