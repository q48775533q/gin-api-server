package model

import (
	"sync"
	"time"
)

// 基础信息，可通用的部分
type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

// 用户信息，用作返回头
type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AssetInfo struct {
	Id        uint64 `json:"id"`
	Assetname string `json:"Assetname"`
	SayHello  string `json:"sayHello"`
	AssetID   string `json:"AssetID"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// 用户列表需要返回的内容,通过锁，对数据进行读取，并且一次性输出。
type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// 用户列表需要返回的内容,通过锁，对数据进行读取，并且一次性输出。
type AssetList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*AssetInfo
}

// 用作返回头使用，限时token
type Token struct {
	Token string `json:"token"`
}
