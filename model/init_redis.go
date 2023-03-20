package model

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func ExampleClient(host string, port string, password string, db int) *redis.Client {

	addr := host + ":" + port
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//fmt.Println("redis_addr:", client)
	return client

}

func InitDefRedis() *redis.Client {
	return ExampleClient(viper.GetString("redis.host"),
		viper.GetString("redis.port"),
		viper.GetString("redis.password"),
		viper.GetInt("redis.db"))

}

func GetREDIS() *redis.Client {
	return InitDefRedis()
}

type REDIS struct {
	Def *redis.Client
}

var RD *REDIS

func (db *REDIS) RInit() {
	RD = &REDIS{
		Def: GetREDIS(),
	}
}
