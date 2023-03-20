package model

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 根据配置文件读取数据库配置
// 在配置文件配置了两个数据源
type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

// 链接数据库
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

// 配置链接池
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// 从参数读取链接参数
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

// 初始化链接信息
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

// 初始化另一组配置，参考最上面的type
func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

// 获得Docker_Mysql具体配置
func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

// 分别初始化数据库GetSelfDB和GetDockerDB
func (db *Database) MInit() {

	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

// 释放数据库链接
func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}
