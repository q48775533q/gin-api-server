package user

import (
	"strconv"

	. "api-server/controller"
	"api-server/model"
	"api-server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 根据id进行删除
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
