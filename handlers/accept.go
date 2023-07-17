package handlers

import (
	"rank/server"
)

type AcceptHandler struct {
}

func (h *AcceptHandler) Handle(cmd *server.Command) {
	cmd.Result = "{\"msg\":\"welcome\"}"
}
