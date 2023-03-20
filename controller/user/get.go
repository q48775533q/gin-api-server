package user

import (
	. "api-server/controller"
	"api-server/model"
	"api-server/pkg/errno"
	"github.com/gin-gonic/gin"
)

// 获得指定用户的信息
func Get(c *gin.Context) {
	username := c.Param("username")

	// log.Info("username:" + username)
	// 从数据库获得指定用户信息
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
