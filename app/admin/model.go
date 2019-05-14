package admin

import (
	"github.com/jinzhu/gorm"
)

// ManageUser table Encryption of user information 用户隐私信息通过 aes 加密处理，不可直观获得
type ManageUser struct {
	gorm.Model
	UUID      string   `gorm:"type:varchar(36);not null;unique; column:uuid"` // 用户 uuid
	Name      string   `gorm:"type:varchar(32)"`                              // 用户名
	NickName  string   `gorm:"type:varchar(64);column:nick_name"`             // 昵称
	RealName  string   `gorm:"type:varchar(64)"`                              // 真实姓名
	Email     string   `gorm:"type:varchar(512);unique"`                      // 邮箱
	Tel       string   `gorm:"type:varchar(512);unique"`                      // 实名制验证
	IDCard    string   `gorm:"type:varchar(512);column:id_card"`              // 证件号码
	CardClass string   `gorm:"type:varchar(512);column:card_class"`           // 证件类型
	Passwd    string   `gorm:"type:varchar(512)"`                             // 密码
	Code      string   `gorm:"type:varchar(32)"`                              // 混淆加密代码
	IP        string   `gorm:"type:varchar(64);column:ip"`                    // 最近一次登陆IP
	Privilege string   `gorm:"type:varchar(1024)"`                            // 拥有的权限
	Audience  []string `gorm:"type:varchar(1024)"`                            // 可访问的域
}

// ManageResets ... 重置密码的作用
type ManageResets struct {
	gorm.Model
	UUID   string `gorm:"type:varchar(36);not null;unique; column:uuid"` // uuid
	Method string `gorm:"type:varchar(32)"`                              // 验证方法
	Email  string `gorm:"type:varchar(512)"`                             // 邮箱
	Tel    string `gorm:"type:varchar(512)"`                             // 实名制验证
	Token  string `gorm:"type:varchar(512)"`                             // 证件类型
}
