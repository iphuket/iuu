package list

import (
	"github.com/gin-gonic/gin"
)

// List Article for List
func List(c *gin.Context) {
	c.Writer.WriteString("Welcome To Article For List")
}
