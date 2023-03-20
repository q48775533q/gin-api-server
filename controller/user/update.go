package user

import (
	"strconv"

	. "api-server/controller"
	"api-server/model"
	"api-server/pkg/errno"
	"api-server/util"

	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"github.com/zxmrlc/log/lager"
)

// 更新用户信息
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// 从url获得用户id
	userId, _ := strconv.Atoi(c.Param("id"))

	// 绑定用户数据
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 根据用户id进行数据更新
	u.Id = uint64(userId)

	// 验证数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密用户密码
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 保存修改的数据
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
