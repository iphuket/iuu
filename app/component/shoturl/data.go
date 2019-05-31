package shoturl

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ShotURL data
type ShotURL struct {
	gorm.Model
	UUID     string `gorm:"type:varchar(36);not null;unique;column:uuid"` // 设置字段为非空并唯一 // `xorm:"varchar(36) pk notnull unique 'uuid'"`
	UserUUID string `gorm:"type:varchar(36); column:user_uuid"`           // `xorm:"varchar(36) pk notnull unique 'user_uuid'"` // user uuid
	Code     string `gorm:"type:varchar(36);unique;column:code"`          // code
	Source   string `gorm:"type:varchar(2048)"`                           // `xorm:"varchar(255) 'source'"`             // 源地址
	Protocol string `gorm:"type:varchar(32)"`
	Exp      time.Time
}
