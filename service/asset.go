package service

import (
	"fmt"
	"github.com/zxmrlc/log"
	"sync"

	"api-server/model"
	"api-server/util"
)

// 查询资产列表的逻辑
func ListAsset(assetname string, offset, limit int) ([]*model.AssetInfo, uint64, error) {
	infos := make([]*model.AssetInfo, 0)
	assets, count, err := model.ListAsset(assetname, offset, limit)
	if err != nil {
		log.Infof("list user err is %s", err)
		return nil, count, err
	}

	var ids []uint64
	for _, asset := range assets {
		ids = append(ids, asset.Id)
		//fmt.Println("ids:", ids)
	}
	wg := sync.WaitGroup{}
	AssetList := model.AssetList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.AssetInfo, len(assetname)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 并行查询
	for _, u := range assets {
		wg.Add(1)
		go func(u *model.AssetModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			AssetList.Lock.Lock()
			defer AssetList.Lock.Unlock()
			AssetList.IdMap[u.Id] = &model.AssetInfo{
				Id:        u.Id,
				Assetname: u.Assetname,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				AssetID:   u.AssetID,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	// go协程等待锁
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, AssetList.IdMap[id])
	}

	return infos, count, nil
}
