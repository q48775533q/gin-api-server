package user

import (
	. "api-server/controller"
	"api-server/model"
	"api-server/pkg/errno"
	"api-server/util"

	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"github.com/zxmrlc/log/lager"
)

// 创建用户
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 验证数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密密码
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// 插入用户到数据库
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// 返回用户信息给用户
	SendResponse(c, nil, rsp)
}
