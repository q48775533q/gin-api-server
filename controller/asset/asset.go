package asset

import "api-server/model"

type CreateRequest struct {
	AssetID   string `json:"assetid"`
	Assetname string `json:"assetname"`
}

type CreateResponse struct {
	Assetname string `json:"assetname"`
}

type ListRequest struct {
	Assetname string `json:"assetname"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64             `json:"totalCount"`
	AssetList  []*model.AssetInfo `json:"userList"`
}
