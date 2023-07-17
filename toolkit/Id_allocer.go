package toolkit

import (
	"sync"
)

type IDAllocer struct {
	counter int64
	mutex   sync.Mutex
	ch      chan int64
}

func (g *IDAllocer) GenerateID(wg *sync.WaitGroup) {
	defer wg.Done()
	g.mutex.Lock()
	for i := 0; i < 125; i++ {
		g.counter++
		id := g.counter
		g.ch <- id // 将生成的ID发送到通道中
	}
	g.mutex.Unlock()
}

func (g *IDAllocer) alloc() (rlt int64) {
	if len(g.ch) < 20 {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go g.GenerateID(wg)
	}
	return <-g.ch
}

var instance *IDAllocer
var once sync.Once

func GetAllocer() *IDAllocer {
	once.Do(func() {
		instance = &IDAllocer{}
		instance.ch = make(chan int64)
		instance.counter = 0
		instance.alloc()
	})
	return instance
}
func Alloc() int64 {
	return GetAllocer().alloc()
}
