package carousel

import (
	"github.com/iphuket/iuu/app/config"
	"fmt"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/iphuket/iuu/app/component/article/model"
)

// Carousel page data
func Carousel(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		carousel, err := get(c.Request.FormValue("class_uuid"), c.Request.FormValue("case_uuid"))
		if err != nil {
			c.JSON(c.Writer.Status(), gin.H{"errCode": "", "errInfo": fmt.Sprint("数据库初始化失败", err)})
			return
		}
		c.JSON(c.Writer.Status(), carousel)
	case "PUT":
		
	}
}

func get(ClassUUID, CaseUUID string) (*model.Carousel, error) {
	carousel := new(model.Carousel)
	db, err := config.DB("mysql")
	if err != nil {
		return nil, err
	}
	db.Where(gin.H{"class_uuid": ClassUUID, "case_uuid": CaseUUID}).Find(&carousel)
	return carousel, nil
}

func put(c *gin.Context) {
	var carousel model.Carousel
	db, err := config.DB("mysql")
	if err != nil {
		c.JSON(c.Writer.Status(), gin.H{"errCode": "", "errInfo": fmt.Sprint("数据库初始化失败", err)})
		return
	}
	carousel.UUID = uuid.New().String() // c.Request.FormValue("c")
	carousel.CaseUUID = c.Request.FormValue("case_uuid")
	carousel.ClassUUID = c.Request.FormValue("class_uuid")
	// carousel.Name =
	db.Create(&carousel)
}

// Delete page data
func Delete(c *gin.Context) {
	// var carousel model.Carousel
	db, err := config.DB("mysql")
	if err != nil {
		c.JSON(c.Writer.Status(), gin.H{"errCode": "", "errInfo": fmt.Sprint("数据库初始化失败", err)})
		return
	}
	db.Where("uuid = ?", c.Request.FormValue("uuid")).Delete(&model.Carousel{})
	c.JSON(c.Writer.Status(), gin.H{"errCode": "success", "uuid": c.Request.FormValue("uuid")})
}
