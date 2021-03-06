package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/pkt/app/component/article/carousel"
	"github.com/iphuket/pkt/app/component/article/list"
)

// Route Settings
func Route(r *gin.RouterGroup) {
	r.Any("carousel", carousel.Carousel)
	r.GET("list", list.List)
}
