package auth

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/config"
	"github.com/iphuket/pkt/library/crypto"
	"github.com/iphuket/pkt/library/passwd"
)

// ResetPassword 开始密码重置 账户 接收 key：重置的对象(email,tel,username) | value：对象内容对应的内容 令牌 token 密码 passwd
func ResetPassword(c *gin.Context) {
	err := verify(c.Request.FormValue("key"), c.Request.FormValue("value"), c.Request.FormValue("token"))
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	// update passwd and code
	var mu ManageUser
	db, err := config.DB()
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	code := crypto.RandomString(6) // 第一次生成 ，第二次用以效果验证
	passwd, err := passwd.New(c.Request.FormValue("passwd"), code)
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	err = db.Model(&mu).Where(c.Request.FormValue("key")+" = ?", c.Request.FormValue("value")).Update(&ManageUser{Passwd: passwd, Code: code}).Error
	if err != nil {
		errorHandle(c, "error", fmt.Sprint(err))
		return
	}
	successHandle(c, "ok")
}

// verify 修改密码验证部分
func verify(key, value, token string) error {
	var mr ManageResets
	db, err := config.DB()
	if err != nil {
		return err
	}
	err = db.Where(key+" = ?", value).Find(&mr).Error
	if err != nil {
		return err
	}
	if token == mr.Token {
		return nil
	}
	return errors.New("token error")
}
