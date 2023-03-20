package user

import (
	. "api-server/controller"
	"api-server/pkg/errno"
	"api-server/service"
	"github.com/gin-gonic/gin"
)

// 根据参数进行查询。支持模糊

/*
curl --location --request GET 'http://127.0.0.1:8080/v1/user?Username=admin&Offset=0&Limit=2' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NzgxMDQ5MjMsImlkIjowLCJuYmYiOjE2NzgxMDQ5MjMsInVzZXJuYW1lIjoiYWRtaW4ifQ.wwqfoqQfnegjyHGRHLkil23GzlCqGuy2RzRq1VrSvgo'
*/
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
