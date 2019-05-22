package shoturl

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/config"
)

// Route Init
func Route(r *gin.RouterGroup) {
	r.Any("l/:key", lctr)
	r.Any("m/:source/:length")
}

// lctr 永久短链接
func lctr(c *gin.Context) {
	db, err := config.DB("mysql")
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	var sURL ShotURL
	// 查询emai
	err = db.Where("key = ?", c.Param(key)).First(&sURL).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	successHandle(c, sURL.Source)
}

// tctr 临时短链接
func tctr(c *gin.Context) {

}
func mctr(c *gin.Context) {

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
