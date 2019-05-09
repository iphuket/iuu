// Package article ... 此包有不少依赖 如 gin,gorm等
// 这个是一个文章管理模块皆在为微信公众号等主题提供文章归类以及阅读。
// 所有文章通过超链接进行链接
// 咱没有计划对内容编辑做出支持
package article

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/iuu/app/component/article/api"
)

// Route Init
func Route(r *gin.RouterGroup) {
	api.Route(r)
}
