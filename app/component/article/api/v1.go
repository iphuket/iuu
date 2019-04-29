package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/iuu/app/component/article/carousel"
	"github.com/iphuket/iuu/app/component/article/list"
)

// Route Settings
func Route(r *gin.RouterGroup) {
	r.GET("home", carousel.Carousel)
	r.GET("list/:class", list.List)
}
