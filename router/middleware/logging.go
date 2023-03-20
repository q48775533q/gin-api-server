package middleware

import (
	"api-server/controller"
	"api-server/model"
	"api-server/pkg/errno"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
	"github.com/zxmrlc/log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 根据请求进行判断，根据uri进行判断
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		// 跳过这几个 requests.
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		//reg := regexp.MustCompile("(/v1/user|/v1/asset|/login)")
		//if !reg.MatchString(path) {
		//	return
		//}

		// 读取body
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// 恢复 io.ReadCloser 到原始状态
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// 基础信息.
		method := c.Request.Method
		ip := c.ClientIP()

		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// 继续往下走.
		c.Next()

		// 计算延时
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		// 本处用作提取用户名，用作日志记录
		var u model.UserModel
		json.Unmarshal(bodyBytes, &u)

		// 获得错误码和message
		var response controller.Response

		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			code = errno.InternalServerError.Code
			message = error.Error(err)
		} else {
			code = response.Code
			message = response.Message
		}

		log.Infof("%-13s | %-12s | %s %s | user: %s | {code: %d, message: %s} ", latency, ip, pad.Right(method, 5, ""), path, u.Username, code, message)
	}
}

func CheckCors() gin.HandlerFunc {
	//这里可以处理一些别的逻辑
	return func(c *gin.Context) {
		// 定义一个origin的map，只有在字典中的key才允许跨域请求
		var allowOrigins = map[string]struct{}{
			"http://127.0.0.1:9527": struct{}{},
		}
		origin := c.Request.Header.Get("Origin") //请求头部
		method := c.Request.Method
		if method == "GET,POST,OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		if origin != "" {
			if _, ok := allowOrigins[origin]; ok {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
				c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
				c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}
		c.Next()
	}
}
