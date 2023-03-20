package controller

import (
	"net/http"

	"api-server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 全局的请求头和返回头
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// 始终返回 http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
