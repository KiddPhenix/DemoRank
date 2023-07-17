package handlers

import (
	"rank/modules"
	"rank/server"
	"strconv"
)

type UpdateRankHandler struct {
}

func (h *UpdateRankHandler) Handle(cmd *server.Command) {
	type resp struct {
		succ int
		msg  string
	}
	if len(cmd.Args) < 2 {
		cmd.Result = "{\"succ\":0, \"msg\":\"args error\"}"

		return
	}
	id, errId := strconv.ParseInt(cmd.Args[0], 10, 0)
	score, errScore := strconv.ParseInt(cmd.Args[1], 10, 0)
	if errId != nil || errScore != nil {
		cmd.Result = "{\"succ\":0, \"msg\":\"args error\"}"
		return
	}
	modules.GetDefaultRank().UpdateItem(id, score)
	cmd.Result = "{\"succ\":1}"
}
