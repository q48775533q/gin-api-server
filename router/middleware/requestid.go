package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求头，存在就使用
		requestId := c.Request.Header.Get("X-Request-Id")

		// 生成uuid4
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// 在应用中使用
		c.Set("X-Request-Id", requestId)

		// 设置请求头，相互传递
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
