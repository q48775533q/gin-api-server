package util

import (
	"api-server/model"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var ctx = context.Background()

var jwt_exp = time.Duration(viper.GetInt64("jwt_exp") * 1000000000)

type RedisStore struct{}

var Rediss RedisStore

// 用作jwt的redis续期
func (rs *RedisStore) JwtSet(key string, val string) error {
	err := model.RD.Def.Set(ctx, key, val, jwt_exp).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// 获取redis里面用户的信息，如果查询不到，返回错误
func (rs *RedisStore) JwtGet(key string) error {
	err := model.RD.Def.Get(ctx, key).Err()
	if err != nil {
		return err
	}
	return err
}
