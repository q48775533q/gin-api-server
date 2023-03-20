package user

import (
	"api-server/model"
)

// 创建的请求头
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 创建的返回头
type CreateResponse struct {
	Username string `json:"username"`
}

// 列表的请求头
type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

// 列表的返回头
type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
