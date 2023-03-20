package model

import (
	"api-server/pkg/constvar"
	"fmt"
	"time"
)

type AssetModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	AssetID   string     `json:"assetid" gorm:"column:assetid;not null" binding:"required"`
	Assetname string     `json:"assetname" gorm:"column:assetname;not null" binding:"required"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

// 数据表
func (c *AssetModel) TableName() string {
	return "tb_asset"
}

// 创建一个新资产.
func (u *AssetModel) Create() error {
	return DB.Self.Create(&u).Error
}

// 查询资产列表
func ListAsset(assetname string, offset, limit int) ([]*AssetModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	assets := make([]*AssetModel, 0)
	var count uint64

	where := fmt.Sprintf("assetname like '%%%s%%'", assetname)

	if err := DB.Self.Model(&AssetModel{}).Where(where).Count(&count).Error; err != nil {
		return assets, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&assets).Error; err != nil {
		return assets, count, err
	}

	//fmt.Println(assets)
	return assets, count, nil

}
