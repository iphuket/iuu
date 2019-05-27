package admin

import (
	"github.com/gin-gonic/gin"
)

// Route Settings
func Route(r *gin.RouterGroup) {
	r.Any("login", Login)
	r.Any("register", Register)
	r.Any("check", Check)
	r.Any("logout", Logout)
	r.Any("reset", Reset)
}
