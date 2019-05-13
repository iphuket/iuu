// Package carousel ...
// 这个包用于处理幻灯片相关
// 技术能力不足，复制话了此包
// 先用再说
package carousel

import (
	"fmt"

	"github.com/iphuket/pkt/app/config"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/component/article/model"
)

// Carousel page data
func Carousel(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		car, err := get(c.Request.FormValue("class_uuid"), c.Request.FormValue("case_uuid"))
		if err != nil {
			errorHandle(c, "error", fmt.Sprint(err))
			return
		}
		successHandle(c, car)

	case "POST":
		do := c.Request.FormValue("do")
		switch do {
		case "UPDATE":
			car := &model.Carousel{UserUUID: "aimo", CaseUUID: c.Request.FormValue("case_uuid"), ClassUUID: c.Request.FormValue("class_uuid"), Name: c.Request.FormValue("name"), Desc: c.Request.FormValue("desc"), Source: c.Request.FormValue("source"), Picture: c.Request.FormValue("picture")}
			err := update(c.Request.FormValue("uuid"), car)
			if err != nil {
				errorHandle(c, "error", fmt.Sprint(err))
				return
			}
			successHandle(c)
		case "CREATE":
			car := &model.Carousel{UUID: uuid.New().String(), UserUUID: "aimo", CaseUUID: c.Request.FormValue("case_uuid"), ClassUUID: c.Request.FormValue("class_uuid"), Name: c.Request.FormValue("name"), Desc: c.Request.FormValue("desc"), Source: c.Request.FormValue("source"), Picture: c.Request.FormValue("picture")}
			car, err := put(car)
			if err != nil {
				errorHandle(c, "error", fmt.Sprint(err))
				return
			}
			successHandle(c, car)
		case "DELETE":
			err := delete(c.Request.FormValue("uuid"))
			if err != nil {
				errorHandle(c, "error", fmt.Sprint(err))
				return
			}
			successHandle(c)
		default:
			errorHandle(c, "error", "not find")
		}
	}

}

func get(ClassUUID, CaseUUID string) (*[]model.Carousel, error) {
	var carousel []model.Carousel
	db, err := config.DB("mysql")
	if err != nil {
		return nil, err
	}

	err = db.Where(&model.Carousel{ClassUUID: ClassUUID, CaseUUID: CaseUUID}).Find(&carousel).Error
	if err != nil {
		return nil, err
	}
	return &carousel, nil
}

func put(c *model.Carousel) (*model.Carousel, error) {
	db, err := config.DB("mysql")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(c)
	// carousel.Name =
	err = db.Create(c).Error
	if err != nil {
		return nil, err
	}
	return c, err
}

// dlete page data Soft Delete
func delete(uuid string) error {
	var carousel model.Carousel
	db, err := config.DB("mysql")
	if err != nil {
		return err
	}
	err = db.Where("uuid = ?", uuid).Delete(carousel).Error
	return err
}

// update page data nothing will be updated as "", 0, false are blank values of their types
func update(uuid string, car *model.Carousel) error {
	var carousel model.Carousel
	db, err := config.DB("mysql")
	if err != nil {
		return err
	}
	err = db.Model(&carousel).Where("uuid = ?", uuid).Updates(car).Error
	return err
}
