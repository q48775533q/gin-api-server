package middleware

import (
	"api-server/controller"
	"api-server/pkg/errno"
	"api-server/pkg/token"
	"github.com/gin-gonic/gin"
)

// 通过token进行鉴权
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析json令牌
		//fmt.Println("auth_print_c:", c)
		if _, err := token.ParseRequest(c); err != nil {
			controller.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		//fmt.Println("AuthMiddleware")
		c.Next()
	}
}
