package token

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	// ErrMissingHeader 标识 `Authorization` 为空.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context Json令牌上下文
type Context struct {
	ID       uint64
	Username string
}

// 验证密钥格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// 确保alg正常可用
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		//fmt.Println("secretFunc")

		return []byte(secret), nil
	}
}

// 使用指定的密钥验证令牌,如果令牌有效，则返回上下文.
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	//fmt.Println("tokenString:", tokenString)
	//fmt.Println("secret:", secret)

	// 分析token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	// Parse error.
	if err != nil {
		return ctx, err

	}

	// 如果令牌有效，读取令牌
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println("claims:", claims)
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)

		//// 打印jwt的超时时间
		//s := claims["exp"].(float64)
		//tm := time.Unix(int64(s), 0)
		//fmt.Println(tm.Sub(time.Now()))

		//// jwt 进行检查，如果redis没有该用户信息，则返回错误，该方式不推荐使用
		//err = util.Rediss.JwtGet(ctx.Username)
		//if err != nil {
		//	return ctx, err
		//}

		//data := []byte(ctx.Username)
		//md5New := md5.New()
		//md5New.Write(data)
		//md5String := hex.EncodeToString(md5New.Sum(nil))
		//
		//// jwt 续期，每次登陆都续期，根据配置文件进行续期 该方式不推荐使用
		//err := util.Rediss.JwtSet(ctx.Username, md5String)
		//if err != nil {
		//	return ctx, nil
		//}
		//fmt.Println("jwt续期成功")

		return ctx, nil

		// 其他错误.
	} else {
		fmt.Println("parse_token3:", token)

		return ctx, err
	}
}

// 分析请求头获得令牌，将其传递给Parse函数进行分析
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// 从配置文件读取 jwt_secret
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// 解析令牌
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// Sign 使用上下文进行签名
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// 读取jwt_secret
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	// 令牌内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + viper.GetInt64("jwt_exp"),
	})

	// 使用指定的密钥对令牌进行签名。
	tokenString, err = token.SignedString([]byte(secret))
	return
}
