package middleware

import (
	"github.com/gin-gonic/gin"
	"gomo/common/runtime"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", runtime.App.GetDb())
	c.Next()
}