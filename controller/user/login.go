package user

import (
	. "api-server/controller"
	"api-server/model"
	"api-server/pkg/auth"
	"api-server/pkg/errno"
	"api-server/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Login 验证登陆并生成 token
/*
curl --location --request POST 'http://127.0.0.1:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"admin","password":"admin"}'
*/
func Login(c *gin.Context) {
	// 监听客户端传来的数据
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 判断用户名是否存在
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, "User:"+u.Username)
		return
	}

	// 对密码进行对比。
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// 验证json token令牌
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")

	fmt.Println("token:", t, err)
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c, nil, model.Token{Token: t})
}
