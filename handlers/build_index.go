package handlers

import (
	"rank/modules"
	"rank/server"
)

type BuildIndexHandler struct {
}

func (h *BuildIndexHandler) Handle(cmd *server.Command) {
	modules.GetDefaultRank().BuildIndex()
	cmd.Result = "{\"succ\":1,\"msg\":\"Index build!\"}"
}
