package modules

import (
	"sync"
)

type playerInfo struct {
	Id int64 //玩家id
	//Name      string    //昵称
	//HeadUrl   string    //头像
	//OtherInfo string    //额外信息
	RankItem *RankItem //
	Index    int
}

type PlayerInfoManager struct {
	dataset map[int64]*playerInfo
}

var instance *PlayerInfoManager
var once sync.Once

func GetPlayerInfoManager() *PlayerInfoManager {
	once.Do(func() {
		instance = &PlayerInfoManager{}
		instance.dataset = make(map[int64]*playerInfo, 1024*1024*64)
	})
	return instance
}

func (mgr *PlayerInfoManager) GetPlayerById(id int64, alloc bool) *playerInfo {
	rlt, isPresent := mgr.dataset[id]
	if !isPresent {
		if alloc {
			rlt = new(playerInfo)
			rlt.Id = id
			rlt.RankItem = nil
			mgr.dataset[id] = rlt
		} else {
			return nil
		}
	}
	return rlt
}

func (mgr *PlayerInfoManager) MutablePlayer(id int64) *playerInfo {
	return mgr.GetPlayerById(id, true)
}

func (mgr *PlayerInfoManager) TryGetPlayer(id int64) *playerInfo {
	return mgr.GetPlayerById(id, false)
}
