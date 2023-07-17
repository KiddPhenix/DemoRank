package handlers

import "rank/server"

type ExitHandler struct {
}

func (h *ExitHandler) Handle(cmd *server.Command) {
	cmd.Result = "{\"succ\":1,\"msg\":\"bye\"}"
	cmd.Session.Quit = true
}
