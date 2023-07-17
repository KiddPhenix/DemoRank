package modules

import (
	"fmt"
	"sync"

	"github.com/huandu/skiplist"
)

type rank struct {
	list         *skiplist.SkipList
	operateIndex int64
}

var onceRank sync.Once
var defaultRank *rank

func GetDefaultRank() *rank {
	onceRank.Do(func() {
		defaultRank = NewRank("default")
	})
	return defaultRank
}

func NewRank(name string) (rlt *rank) {
	rlt = new(rank)
	rlt.list = skiplist.New(skiplist.GreaterThanFunc(func(k1, k2 interface{}) int {

		s1 := k1.(*RankItem).Score
		s2 := k2.(*RankItem).Score
		if s1 == s2 {
			t1 := k1.(*RankItem).Time
			t2 := k2.(*RankItem).Time
			return int(t1 - t2)

		}
		return int(s2 - s1)
	}))
	rlt.list.SetMaxLevel(24)
	return
}

func (r *rank) UpdateItem(id int64, score int64) {
	r.operateIndex += 1
	player := GetPlayerInfoManager().MutablePlayer(id)
	if player.RankItem != nil {
		elem := r.list.Get(player.RankItem)
		if elem != nil {
			r.list.Remove(player.RankItem)
		}
	} else {
		player.RankItem = new(RankItem)
		player.RankItem.Score = 0
	}
	player.RankItem.Id = id
	if score > player.RankItem.Score {
		player.RankItem.Score = score
		player.RankItem.Time = r.operateIndex
	}
	player.RankItem.Player = player
	r.list.Set(player.RankItem, player.RankItem)
}

func (r *rank) BuildIndex() {
	//e := r.list.Front()
	index := 1
	for e := r.list.Front(); e != nil; e = e.Next() {
		v := e.Value.(*RankItem)
		v.Player.Index = index
		index += 1
	}
}

func (r *rank) GetRankItemElem(id int64) *skiplist.Element {
	player := GetPlayerInfoManager().TryGetPlayer(id)
	if player != nil && player.RankItem != nil {
		elem := r.list.Get(player.RankItem)
		return elem
	}
	return nil
}
func (r *rank) GetLastOne() *skiplist.Element {
	return r.list.Back()
}

////////////////////////////////////// dump to console /////////////////////////////////////

func (r *rank) DumpItem(id int64) {
	player := GetPlayerInfoManager().TryGetPlayer(id)
	if player != nil && player.RankItem != nil {
		elem := r.list.Get(player.RankItem)
		fmt.Print("*  ")
		fmt.Println(elem.Value)
		counter := 5
		for e := elem.Next(); e != nil; e = e.Next() {
			v := e.Value
			fmt.Print("-  ")
			fmt.Println(v)
			counter -= 1
			if counter <= 0 {
				break
			}
		}
		counter = 5
		for e := elem.Prev(); e != nil; e = e.Prev() {
			v := e.Value
			fmt.Print("+  ")
			fmt.Println(v)
			counter -= 1
			if counter <= 0 {
				break
			}
		}
	}
}

func (r *rank) PrintRank() {
	fmt.Println(r.list.Len())
	counter := 10
	for e := r.list.Front(); e != nil; e = e.Next() {
		v := e.Value
		fmt.Println(v)
		counter -= 1
		if counter <= 0 {
			break
		}
	}
	fmt.Println("-------")
	counter = 10
	for e := r.list.Back(); e != nil; e = e.Prev() {
		v := e.Value
		fmt.Println(v)
		counter -= 1
		if counter <= 0 {
			return
		}
	}
}
