package asset

import (
	. "api-server/controller"
	"api-server/model"
	"api-server/pkg/errno"
	"api-server/util"
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"github.com/zxmrlc/log/lager"
)

// 创建资产
func Create(c *gin.Context) {
	log.Info("Asset Create function called. ", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.AssetModel{
		AssetID:   r.AssetID,
		Assetname: r.Assetname,
	}

	// 插入用户到数据库
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Assetname: r.Assetname,
	}

	// 返回用户信息给用户
	SendResponse(c, nil, rsp)
}
