package test

import (
	"math/rand"
	Rank "rank/modules"
	tk "rank/toolkit"
)

func TestOnce() {
	tk.L("Test started !")
	rank := Rank.NewRank("Tester")
	rank.UpdateItem(1, 100)
	rank.UpdateItem(2, 200)
	rank.UpdateItem(3, 200)
	rank.UpdateItem(1, 100500)
	nMax := (int64)(10000000)
	//nMax := (int64)(100000)
	var i int64
	var counter int64 = 0
	for i = 0; i < nMax*5; i++ {
		id := i%nMax + 1
		score := rand.Int63() % 1000000
		rank.UpdateItem(id, score)
		counter += 1
		if counter >= nMax {
			counter = 0
			tk.L("Looped once ")
		}
	}
	tk.L("set data finished!")
	rank.BuildIndex()
	tk.L("Build Index finished!")
	rank.DumpItem(8888)
	rank.PrintRank()
}

func FillTestDataToDefaultRank() {
	rank := Rank.GetDefaultRank()
	nMax := (int64)(100)
	//nMax := (int64)(100000)
	var i int64
	for i = 0; i < nMax; i++ {
		id := i%nMax + 1
		score := rand.Int63() % 1000000
		rank.UpdateItem(id, score)
		tk.L("Fill test data: id:", id, " score:", score)
	}
}
