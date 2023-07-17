package server

type UnknownHandler struct {
}

func (h *UnknownHandler) Handle(cmd *Command) {
	cmd.Result = "{\"succ\":0,\"msg\":\"invalid cmd!\"}"
}

func CreateChannel() chan Command {
	return make(chan Command, 1024)
}

func LoopEvent(ch chan Command, router map[string]Handler) {
	var exitFlag bool = false
	for !exitFlag {
		cmd, ok := <-ch
		if ok {
			handler, exist := router[cmd.Cmd]
			if !exist {
				handler = &UnknownHandler{}
				exist = true
			}
			if exist {
				handler.Handle(&cmd)
				cmd.Conn.Write([]byte(cmd.Result))
				cmd.Conn.Write([]byte("\r\n"))
				if cmd.Session.Quit {
					cmd.Session.Alive = false
				}
			}
		}
	}
}
