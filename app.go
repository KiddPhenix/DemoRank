package main

import (
	"rank/handlers"
	"rank/server"
	"rank/test"
	tk "rank/toolkit"
)

func main() {
	test.FillTestDataToDefaultRank()
	tk.L("Hello, Rank server started!")
	srv := server.InitServer("tcp", 23)
	srv.RegHandler("_init", &handlers.AcceptHandler{})
	srv.RegHandler("update", &handlers.UpdateRankHandler{})
	srv.RegHandler("exit", &handlers.ExitHandler{})
	srv.RegHandler("rank", &handlers.GetSubRankHandler{})
	srv.RegHandler("build", &handlers.BuildIndexHandler{})
	srv.LoopForever()
	return
}
