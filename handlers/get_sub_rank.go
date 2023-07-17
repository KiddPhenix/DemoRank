package handlers

import (
	"fmt"
	"rank/modules"
	"rank/server"
	"strconv"
)

type GetSubRankHandler struct {
}

func (h *GetSubRankHandler) Handle(cmd *server.Command) {
	const rankUpperSize int = 10
	const rankFillSize int = 10 + 10 + 1
	if len(cmd.Args) < 1 {
		cmd.Result = "{\"succ\":0, \"msg\":\"args error: too few args , should be < rank,id >\"}"
		return
	}
	id, errId := strconv.ParseInt(cmd.Args[0], 10, 0)
	if errId != nil {
		cmd.Result = "{\"succ\":0, \"msg\":\"args error: id format invalid\"}"
		return
	}
	//find sub rank head
	elem := modules.GetDefaultRank().GetRankItemElem(id)
	if elem == nil {
		elem = modules.GetDefaultRank().GetLastOne()
	}

	for cnt := 0; cnt < rankUpperSize && elem.Prev() != nil; elem = elem.Prev() {
		cnt += 1
	}

	//fill sub rank
	var resultCnt int = 0
	var result [rankFillSize]*modules.RankItem
	for resultCnt = 0; resultCnt < rankFillSize && elem != nil; elem = elem.Next() {
		result[resultCnt] = elem.Value.(*modules.RankItem)
		resultCnt += 1
	}
	//fill to msg
	cmd.Result = "{\r\n\"rank\":[\r\n"
	for i := 0; i < resultCnt; i++ {
		index := -1
		rankItem := result[i]
		if rankItem.Player != nil {
			index = rankItem.Player.Index
		}
		if i > 0 {
			cmd.Result += ",\r\n"
		}
		cmd.Result += "\t"
		cmd.Result += fmt.Sprintf("{\"id\": %d,\"score\": %d, \"index\":%d}",
			rankItem.Id,
			rankItem.Score,
			index)

	}
	cmd.Result += "\r\n],\r\n\"succ\":1}"
}
