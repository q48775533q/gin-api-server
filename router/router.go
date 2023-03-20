package router

import (
	"net/http"

	"api-server/controller/asset"
	"api-server/controller/sd"
	"api-server/controller/user"
	"api-server/router/middleware"

	"github.com/gin-gonic/gin"
)

// 读取中间件、路由，controller
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	// 处理异常的
	g.Use(gin.Recovery())

	// 处理客户端异常。
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	// 载入自定义内容
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 用作身份验证
	g.POST("/v1/login", user.Login)

	// 需要身份验证才能操作的api
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("", user.List)
		u.GET("/:username", user.Get)
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
	}

	a := g.Group("/v1/asset")
	a.Use(middleware.AuthMiddleware())
	{
		a.GET("", asset.List)
		a.POST("/:id", asset.Create)
	}

	// 健康检查程序
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
