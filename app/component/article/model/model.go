package model

import (
	"github.com/jinzhu/gorm"

)

// Carousel sql data struct
type Carousel struct {
	gorm.Model
	UUID      string `gorm:"type:varchar(36);not null;unique; column:uuid"` // 设置字段为非空并唯一 // `xorm:"varchar(36) pk notnull unique 'uuid'"`      // uuid
	UserUUID  string `gorm:"type:varchar(36); column:user_uuid"`            // `xorm:"varchar(36) pk notnull unique 'user_uuid'"` // user uuid
	CaseUUID  string `gorm:"type:varchar(36); column:case_uuid"`            // `xorm:"varchar(64) 'class_uuid'"`          // case_uuid
	ClassUUID string `gorm:"type:varchar(36); column:class_uuid"`           // `xorm:"varchar(64) 'class_uuid'"`                  // 上级类目
	Name      string `gorm:"type:varchar(255)"`                             // `xorm:"varchar(255) 'name'"`               // 名称
	Desc      string `gorm:"type:varchar(255)"`                             // `xorm:"varchar(255) 'desc'"`               // 摘要
	Picture   string `gorm:"type:varchar(1024)"`                            //`xorm:"varchar(255) 'picture'"`            // 图像url 地址
	Source    string `gorm:"type:varchar(1024)"`                            // `xorm:"varchar(255) 'source'"`             // 源地址 文章地址

}

// List for data struct
type List struct {
	Article []Article
}

// Article data struct
type Article struct {
	gorm.Model
	UUID      string `gorm:"type:varchar(36);not null;unique; column:uuid"` // uuid
	UserUUID  string `gorm:"type:varchar(36); column:user_uuid"`            //`xorm:"varchar(36) pk notnull unique 'user_uuid'"` // user uuid
	CaseUUID  string `gorm:"type:varchar(36); column:case_uuid"`            // `xorm:"varchar(64) 'case_uuid'"`             // case_uuid
	ClassUUID string `gorm:"type:varchar(36); column:class_uuid"`           // `xorm:"varchar(64) 'class_uuid'"`                  // 上级类目
	Name      string `gorm:"type:varchar(255)"`                             // `xorm:"varchar(255) 'name'"`                  // 名称
	Desc      string `gorm:"type:varchar(255)"`                             // `xorm:"varchar(255) 'desc'"`                  // 摘要
	Content   string `gorm:"type:text"`                                     // `xorm:"varchar(255) 'content'"`               // 主要内容
}

// Class for Article
type Class struct {
	gorm.Model
	UUID      string `gorm:"type:varchar(36);not null;unique; column:uuid"` // uuid
	UserUUID  string `gorm:"type:varchar(36); column:user_uuid"`            //`xorm:"varchar(36) pk notnull unique 'user_uuid'"` // user uuid
	CaseUUID  string `gorm:"type:varchar(36); column:case_uuid"`            // case_uuid
	ClassUUID string `gorm:"type:varchar(36); column:class_uuid"`           // `xorm:"varchar(64) 'class_uuid'"`                  // 上级类目
	Level     int64  `gorm:"type:int"`                                      // 分类级别
	Name      string `gorm:"varchar(64) 'name'"`                            // 名称

}

// Case for article men
type Case struct {
	gorm.Model
	UUID     string `gorm:"type:varchar(36);not null;unique; column:uuid"`
	UserUUID string `gorm:"type:varchar(36); column:user_uuid"` // user uuid

}

func db() (*gorm.DB, error) {
	return nil,nil
}

// GetCase ...
func GetCase(UserUUID string) (*Case, error) {
	return nil,nil
}

// GetClass ...
func GetClass(UserUUID, CaseUUID string) *Class {
	return nil
}

// GetArticle ...
func GetArticle(UserUUID, CaseUUID, ClassUUID string) *Article {
	return nil
}

// GetCarousel ...
func GetCarousel(UserUUID, CaseUUID string) *Carousel {
	return nil
}
