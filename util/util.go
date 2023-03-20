package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// 生成随机ID
func GenShortId() (string, error) {
	return shortid.Generate()
}

// 获得请求头的X-Request-Id
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
